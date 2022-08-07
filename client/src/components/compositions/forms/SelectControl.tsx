import React, { ChangeEvent, FocusEvent, ReactElement, ReactNode } from 'react';
import Feedback from '@components/blocks/forms/Feedback';
import FormGroup from '@components/blocks/forms/FormGroup';
import FormLabel from '@components/blocks/forms/FormLabel';
import FormText from '@components/blocks/forms/FormText';
import Select from '@components/blocks/forms/Select';

export type SelectControlProps = {
  children?: ReactNode;
  id?: string;
  name?: string;
  className?: string;
  defaultValue?: string | string[] | number;
  value?: string | string[] | number;
  label?: string;
  placeholder?: string;
  tabIndex?: number;
  error?: string;
  tip?: string;
  selectId?: string;
  labelId?: string;
  errorId?: string;
  tipId?: string;
  isValid?: boolean;
  required?: boolean;
  disabled?: boolean;
  theme?: {
    select?: string;
    label?: string;
    tip?: string;
    error?: string;
  };
  onChange?: (event: ChangeEvent<HTMLSelectElement>) => void;
  onBlur?: (event: FocusEvent<HTMLSelectElement>) => void;
};

const SelectControl = (props: SelectControlProps): ReactElement => {
  const {
    children,
    id,
    name,
    className,
    defaultValue,
    value,
    label,
    placeholder,
    tabIndex,
    error,
    tip,
    selectId,
    labelId,
    errorId,
    tipId,
    isValid,
    required,
    disabled,
    theme = {},
    onChange,
    onBlur
  } = props;

  return (
    <FormGroup id={id} style="default" className={className}>
      {label && (
        <FormLabel id={labelId} className={theme.label} style="default" htmlFor={selectId} required={required}>
          {label}
        </FormLabel>
      )}
      <Select
        id={selectId}
        name={name}
        className={theme.select}
        defaultValue={defaultValue}
        value={value}
        placeholder={placeholder}
        tabIndex={tabIndex}
        isValid={isValid}
        required={required}
        disabled={disabled}
        onChange={onChange}
        onBlur={onBlur}>
        {children}
      </Select>
      {error && (
        <Feedback id={errorId} className={theme.error} isValid={isValid}>
          {error}
        </Feedback>
      )}
      {!error && tip && (
        <FormText id={tipId} className={theme.tip} block>
          {tip}
        </FormText>
      )}
    </FormGroup>
  );
};

export default SelectControl;
