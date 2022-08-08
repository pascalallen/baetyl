import jwtDecode from 'jwt-decode';
import moment from 'moment';
import { User } from '@domain/data/User';
import ContextStore from '@stores/ContextStore';
import SessionStore from '@stores/SessionStore';
import auth from '@services/_api/api/auth';
import { listToMap } from '@services/utilities/collections';
import { RegisterFormValues } from '@pages/register/RegisterPage';

class SecurityContext {
  private readonly sessionStore: SessionStore;
  private readonly contextStore: ContextStore;

  constructor(sessionStore: SessionStore, contextStore: ContextStore) {
    this.sessionStore = sessionStore;
    this.contextStore = contextStore;
  }

  public async register(params: RegisterFormValues): Promise<void> {
    const response = await auth.register(params);
    this.sessionStore.setSession({
      token: response.body.data?.token,
      data: jwtDecode(response.body.data?.token),
      clientIat: moment().unix(),
      roles: listToMap(response.body.data?.roles),
      permissions: listToMap(response.body.data?.permissions)
    });
    this.contextStore.setUser(response.body.data.user);
  }

  public isLoggedIn(): boolean {
    return this.sessionStore.hasSession();
  }

  public getUser(): User | undefined {
    return this.contextStore.getUser();
  }

  public hasPermission(permission: string): boolean {
    if (!this.sessionStore.hasSession()) {
      return false;
    }

    return this.sessionStore.getSession()?.permissions[permission] === permission;
  }

  public hasPermissions(permissions: string[]): boolean {
    if (!this.sessionStore.hasSession()) {
      return false;
    }

    return permissions.every((permission: string) => {
      return this.sessionStore.getSession()?.permissions[permission] === permission;
    });
  }

  public hasRole(role: string): boolean {
    if (!this.sessionStore.hasSession()) {
      return false;
    }

    return this.sessionStore.getSession()?.roles[role] === role;
  }
}

export default SecurityContext;
