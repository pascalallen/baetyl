import useStore from '@hooks/useStore';
import SecurityContext from '@services/security/SecurityContext';

const useSecurity = (): SecurityContext => {
  const sessionStore = useStore('sessionStore');
  const contextStore = useStore('contextStore');

  return new SecurityContext(sessionStore, contextStore);
};

export default useSecurity;
