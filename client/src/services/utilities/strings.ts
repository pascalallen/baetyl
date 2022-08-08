import moment from 'moment';
import { Json } from '@domain/types/Json';

export const queryStringify = (params: { [key: string]: Json }): string => {
  const output: string[] = [];

  Object.keys(params).forEach(key => {
    const value = params[key];
    if (Array.isArray(value)) {
      output.push(`${key}[]=${value.join(`&${key}[]=`)}`);
    } else {
      output.push(`${key}=${value}`);
    }
  });

  return output.length > 0 ? `?${output.join('&')}` : '';
};

export const formatDateTime = (dateString: string): string => {
  return moment(dateString).format('L');
};
