import React, { FocusEvent, ChangeEvent, ReactElement, useState } from 'react';
import { useFormik } from 'formik';
import { useNavigate } from 'react-router';
import Path from '@domain/constants/Path';
import useSecurity from '@hooks/useSecurity';
import { isApiError, isApiFail } from '@services/_api/ApiService';
import notifications from '@services/notifications/notifications';
import form from '@services/utilities/form';
import Button from '@components/blocks/buttons/Button';
import Form from '@components/blocks/forms/Form';
import InputTextControl from '@components/compositions/forms/InputTextControl';

export type RegisterFormValues = {
  first_name: string;
  last_name: string;
  email_address: string;
  password: string;
  confirm_password: string;
};

const initialFormValues: RegisterFormValues = {
  first_name: '',
  last_name: '',
  email_address: '',
  password: '',
  confirm_password: ''
};

type ServerErrors = {
  first_name: string;
  last_name: string;
  email_address: string;
  password: string;
  confirm_password: string;
};

type State = {
  serverErrors: ServerErrors;
};

const initialState: State = {
  serverErrors: {
    first_name: '',
    last_name: '',
    email_address: '',
    password: '',
    confirm_password: ''
  }
};

const RegisterPage = (): ReactElement => {
  const navigate = useNavigate();
  const securityContext = useSecurity();

  const [serverErrors, setServerErrors] = useState(initialState.serverErrors);

  const handleServerValidationErrors = (errors: ServerErrors): void => {
    setServerErrors({
      ...initialState.serverErrors,
      first_name: errors.first_name || '',
      last_name: errors.last_name || '',
      email_address: errors.email_address || '',
      password: errors.password || '',
      confirm_password: errors.confirm_password || ''
    });
  };

  const handleSubmit = async (formData: RegisterFormValues): Promise<void> => {
    await securityContext
      .register(formData)
      .catch((error: unknown) => {
        if (isApiFail(error)) {
          return handleServerValidationErrors(error.body.data as ServerErrors);
        }

        if (isApiError(error)) {
          notifications.error(error.body.message);
        }
      })
      .finally(() => {
        formik.setSubmitting(false);
      });
    navigate(Path.REGISTER);
  };

  const validate = (formData: RegisterFormValues): void => console.log(formData);

  const formik = useFormik({
    initialValues: initialFormValues,
    validate: validate,
    onSubmit: handleSubmit,
    validateOnBlur: true,
    validateOnChange: false
  });

  const handleChange = (event: ChangeEvent<HTMLInputElement>): void => formik.handleChange(event);

  const handleBlur = (event: FocusEvent<HTMLInputElement>): void => {
    setServerErrors(initialState.serverErrors);
    formik.handleBlur(event);
  };

  const firstNameInputProps = form.getInputProps('first_name', formik, serverErrors);
  const lastNameInputProps = form.getInputProps('last_name', formik, serverErrors);
  const emailAddressInputProps = form.getInputProps('email_address', formik, serverErrors);
  const passwordInputProps = form.getInputProps('password', formik, serverErrors);
  const confirmPasswordInputProps = form.getInputProps('confirm_password', formik, serverErrors);

  return (
    <div className="container">
      <div className="row">
        <div className="col">
          <Form id="register-form" className="register-form">
            <InputTextControl
              inputId="register-form-first-name"
              className="register-form-first-name mb-3"
              name="first_name"
              type="text"
              label="First name"
              tabIndex={1}
              value={formik.values.first_name}
              isValid={firstNameInputProps.isValid}
              error={firstNameInputProps.errorMessage}
              required
              onChange={handleChange}
              onBlur={handleBlur}
            />
            <InputTextControl
              inputId="register-form-last-name"
              className="register-form-last-name mb-3"
              name="last_name"
              type="text"
              label="Last name"
              tabIndex={2}
              value={formik.values.last_name}
              isValid={lastNameInputProps.isValid}
              error={lastNameInputProps.errorMessage}
              required
              onChange={handleChange}
              onBlur={handleBlur}
            />
            <InputTextControl
              inputId="register-form-email-address"
              className="register-form-email-address mb-3"
              name="email_address"
              type="email"
              label="Email address"
              tabIndex={3}
              value={formik.values.email_address}
              isValid={emailAddressInputProps.isValid}
              error={emailAddressInputProps.errorMessage}
              required
              onChange={handleChange}
              onBlur={handleBlur}
            />
            <InputTextControl
              inputId="register-form-password"
              className="register-form-password mb-3"
              name="password"
              type="password"
              label="Password"
              tabIndex={4}
              value={formik.values.password}
              isValid={passwordInputProps.isValid}
              error={passwordInputProps.errorMessage}
              required
              onChange={handleChange}
              onBlur={handleBlur}
            />
            <InputTextControl
              inputId="register-form-confirm-password"
              className="register-form-confirm-password mb-3"
              name="confirm_password"
              type="password"
              label="Confirm password"
              tabIndex={5}
              value={formik.values.confirm_password}
              isValid={confirmPasswordInputProps.isValid}
              error={confirmPasswordInputProps.errorMessage}
              required
              onChange={handleChange}
              onBlur={handleBlur}
            />
            <Button
              id="register-form-submit"
              className="register-form-submit"
              type="submit"
              tabIndex={6}
              disabled={formik.isSubmitting}>
              Sign In
            </Button>
          </Form>
        </div>
      </div>
    </div>
  );
};

export default RegisterPage;
