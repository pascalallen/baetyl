import React, { ReactElement, ReactNode } from 'react';

export type OptionProps = {
  children?: ReactNode;
  id?: string;
  className?: string;
  value?: string | string[] | number;
  disabled?: boolean;
};

const Option = (props: OptionProps): ReactElement => {
  const { children, id, className, value, disabled } = props;

  return (
    <option id={id} className={className} value={value} disabled={disabled}>
      {children}
    </option>
  );
};

export default Option;
