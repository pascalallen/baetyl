import React, { MouseEvent, ReactElement, ReactNode } from 'react';
import classNames from 'classnames';
import { StyleButtonVariants } from '@domain/constants/StyleVariants';

export type ButtonProps = {
  children?: ReactNode;
  id?: string;
  className?: string;
  type?: 'submit' | 'reset' | 'button';
  tabIndex?: number;
  disabled?: boolean;
  variant?: StyleButtonVariants;
  size?: 'sm' | 'lg';
  onClick?: (event?: MouseEvent<HTMLButtonElement>) => void;
};

const Button = (props: ButtonProps): ReactElement => {
  const { children, id, className, type = 'button', tabIndex, disabled, variant = 'primary', size, onClick } = props;

  const classes: string[] = ['btn'];

  if (variant) {
    classes.push(`btn-${variant}`);
  }

  if (size) {
    classes.push(`btn-${size}`);
  }

  return (
    <button
      id={id}
      className={classNames(...classes, className)}
      tabIndex={tabIndex}
      disabled={disabled}
      type={type}
      onClick={onClick}>
      {children}
    </button>
  );
};

export default Button;
