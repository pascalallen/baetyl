import ContextStore from '@stores/ContextStore';
import SessionStore from '@stores/SessionStore';

class RootStore {
  public sessionStore: SessionStore;
  public contextStore: ContextStore;

  constructor() {
    this.sessionStore = new SessionStore();
    this.contextStore = new ContextStore();
  }
}

export default RootStore;
