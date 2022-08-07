import React, { ReactElement, ChangeEvent, FocusEvent } from 'react';
import Feedback from '@components/blocks/forms/Feedback';
import FormGroup from '@components/blocks/forms/FormGroup';
import FormLabel from '@components/blocks/forms/FormLabel';
import FormText from '@components/blocks/forms/FormText';
import Radio from '@components/blocks/forms/Radio';

export type RadioControlProps = {
  id?: string;
  name?: string;
  className?: string;
  formName?: string;
  value?: string | string[] | number;
  label?: string;
  tabIndex?: number;
  error?: string;
  tip?: string;
  inputId?: string;
  labelId?: string;
  errorId?: string;
  tipId?: string;
  checked?: boolean;
  inline?: boolean;
  isValid?: boolean;
  required?: boolean;
  disabled?: boolean;
  theme?: {
    input?: string;
    label?: string;
    tip?: string;
    error?: string;
  };
  onChange?: (event: ChangeEvent<HTMLInputElement>) => void;
  onBlur?: (event: FocusEvent<HTMLInputElement>) => void;
};

const RadioControl = (props: RadioControlProps): ReactElement => {
  const {
    id,
    name,
    className,
    value,
    label,
    tabIndex,
    error,
    tip,
    inputId,
    labelId,
    errorId,
    tipId,
    checked,
    inline,
    isValid,
    required,
    disabled,
    theme = {},
    onChange,
    onBlur
  } = props;

  return (
    <FormGroup id={id} style="radio" className={className} inline={inline}>
      <Radio
        id={inputId}
        name={name}
        className={theme.input}
        value={value}
        tabIndex={tabIndex}
        isValid={isValid}
        checked={checked}
        required={required}
        disabled={disabled}
        onChange={onChange}
        onBlur={onBlur}
      />
      <FormLabel id={labelId} className={theme.label} style="radio" htmlFor={inputId} required={required}>
        {label}
      </FormLabel>
      {error && (
        <Feedback id={errorId} className={theme.error} isValid={false}>
          {error}
        </Feedback>
      )}
      {!error && tip && (
        <FormText id={tipId} className={theme.tip}>
          {tip}
        </FormText>
      )}
    </FormGroup>
  );
};

export default RadioControl;
