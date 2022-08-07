import React, { ReactElement, ReactNode } from 'react';
import classNames from 'classnames';

export type ButtonToolbarProps = {
  screenReaderLabel: string;
  children?: ReactNode;
  id?: string;
  className?: string;
};

const ButtonToolbar = (props: ButtonToolbarProps): ReactElement => {
  const { screenReaderLabel, children, id, className } = props;

  return (
    <div id={id} className={classNames('btn-toolbar', className)} role="toolbar" aria-label={screenReaderLabel}>
      {children}
    </div>
  );
};

export default ButtonToolbar;
