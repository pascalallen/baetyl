import { makeAutoObservable } from 'mobx';
import { Session } from '@domain/data/Session';

class SessionStore {
  private session?: Session;

  constructor() {
    makeAutoObservable(this);
  }

  public setSession(session: Session): void {
    this.session = Object.freeze(session);
  }

  public clearSession(): void {
    delete this.session;
  }

  public hasSession(): boolean {
    return this.session !== undefined;
  }

  public getSession(): Session | undefined {
    return this.session;
  }
}

export default SessionStore;
