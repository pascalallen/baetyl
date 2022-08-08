import { makeAutoObservable } from 'mobx';
import { User } from '@domain/data/User';

class ContextStore {
  private user?: User;

  constructor() {
    makeAutoObservable(this);
  }

  public setUser(user: User): void {
    this.user = Object.freeze(user);
  }

  public clearUser(): void {
    delete this.user;
  }

  public hasUser(): boolean {
    return this.user !== undefined;
  }

  public getUser(): User | undefined {
    return this.user;
  }
}

export default ContextStore;
