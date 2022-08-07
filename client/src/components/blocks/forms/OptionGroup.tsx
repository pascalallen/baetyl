import React, { ReactElement, ReactNode } from 'react';

export type OptionGroupProps = {
  children?: ReactNode;
  id?: string;
  className?: string;
  label?: string;
  disabled?: boolean;
};

const OptionGroup = (props: OptionGroupProps): ReactElement => {
  const { children, id, className, label, disabled } = props;

  return (
    <optgroup id={id} className={className} label={label} disabled={disabled}>
      {children}
    </optgroup>
  );
};

export default OptionGroup;
