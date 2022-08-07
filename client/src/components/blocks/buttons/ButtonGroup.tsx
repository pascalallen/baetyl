import React, { ReactElement, ReactNode } from 'react';
import classNames from 'classnames';

export type ButtonGroupProps = {
  screenReaderLabel: string;
  children?: ReactNode;
  id?: string;
  className?: string;
  size?: 'sm' | 'lg';
  vertical?: boolean;
};

const ButtonGroup = (props: ButtonGroupProps): ReactElement => {
  const { screenReaderLabel, children, id, className, size, vertical } = props;

  const classes: string[] = [];

  if (size) {
    classes.push(`btn-group-${size}`);
  }

  if (vertical) {
    classes.push('btn-group-vertical');
  } else {
    classes.push('btn-group');
  }

  return (
    <div id={id} className={classNames(...classes, className)} role="group" aria-label={screenReaderLabel}>
      {children}
    </div>
  );
};

export default ButtonGroup;
