import { DomainEvents } from '@domain/constants/DomainEvents';
import eventDispatcher from '@services/events/eventDispatcher';

const defaultHeader = 'Notification';

class NotificationService {
  public primary(message: string, header?: string): void {
    eventDispatcher.dispatch({
      name: DomainEvents.NOTIFICATION,
      data: {
        type: 'primary',
        header: header || defaultHeader,
        message: message
      }
    });
  }

  public secondary(message: string, header?: string): void {
    eventDispatcher.dispatch({
      name: DomainEvents.NOTIFICATION,
      data: {
        type: 'secondary',
        header: header || defaultHeader,
        message: message
      }
    });
  }

  public success(message: string, header?: string): void {
    eventDispatcher.dispatch({
      name: DomainEvents.NOTIFICATION,
      data: {
        type: 'success',
        header: header || defaultHeader,
        message: message
      }
    });
  }

  public error(message: string, header?: string): void {
    eventDispatcher.dispatch({
      name: DomainEvents.NOTIFICATION,
      data: {
        type: 'danger',
        header: header || defaultHeader,
        message: message
      }
    });
  }

  public warning(message: string, header?: string): void {
    eventDispatcher.dispatch({
      name: DomainEvents.NOTIFICATION,
      data: {
        type: 'warning',
        header: header || defaultHeader,
        message: message
      }
    });
  }

  public info(message: string, header?: string): void {
    eventDispatcher.dispatch({
      name: DomainEvents.NOTIFICATION,
      data: {
        type: 'info',
        header: header || defaultHeader,
        message: message
      }
    });
  }

  public dark(message: string, header?: string): void {
    eventDispatcher.dispatch({
      name: DomainEvents.NOTIFICATION,
      data: {
        type: 'dark',
        header: header || defaultHeader,
        message: message
      }
    });
  }

  public light(message: string, header?: string): void {
    eventDispatcher.dispatch({
      name: DomainEvents.NOTIFICATION,
      data: {
        type: 'light',
        header: header || defaultHeader,
        message: message
      }
    });
  }
}

export default NotificationService;
