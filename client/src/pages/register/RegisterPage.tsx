import React, { ReactElement } from 'react';
import { useFormik } from 'formik';
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

const RegisterPage = (): ReactElement => {
  const handleSubmit = (formData: RegisterFormValues): void => {
    console.log(formData);
  };

  const formik = useFormik({
    initialValues: initialFormValues,
    validate: () => console.log('TODO'),
    onSubmit: handleSubmit,
    validateOnBlur: true,
    validateOnChange: false
  });

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
              // isValid={firstNameInputProps.isValid}
              // required={firstNameInputProps.isRequired}
              // error={firstNameInputProps.errorMessage}
              // onChange={handleChange}
              // onBlur={handleBlur}
            />
            <InputTextControl
              inputId="register-form-last-name"
              className="register-form-last-name mb-3"
              name="last_name"
              type="text"
              label="Last name"
              tabIndex={2}
              value={formik.values.last_name}
              // isValid={lastNameInputProps.isValid}
              // required={lastNameInputProps.isRequired}
              // error={lastNameInputProps.errorMessage}
              // onChange={handleChange}
              // onBlur={handleBlur}
            />
            <InputTextControl
              inputId="register-form-email-address"
              className="register-form-email-address mb-3"
              name="email_address"
              type="email"
              label="Email address"
              tabIndex={3}
              value={formik.values.email_address}
              // isValid={emailInputProps.isValid}
              // required={emailInputProps.isRequired}
              // error={emailInputProps.errorMessage}
              // onChange={handleChange}
              // onBlur={handleBlur}
            />
            <InputTextControl
              inputId="register-form-password"
              className="register-form-password mb-3"
              name="password"
              type="password"
              label="Password"
              tabIndex={4}
              value={formik.values.password}
              // isValid={passwordInputProps.isValid}
              // required={passwordInputProps.isRequired}
              // error={passwordInputProps.errorMessage}
              // onChange={handleChange}
              // onBlur={handleBlur}
            />
            <InputTextControl
              inputId="register-form-confirm-password"
              className="register-form-confirm-password mb-3"
              name="confirm_password"
              type="password"
              label="Confirm password"
              tabIndex={5}
              value={formik.values.confirm_password}
              // isValid={confirmPasswordInputProps.isValid}
              // required={confirmPasswordInputProps.isRequired}
              // error={confirmPasswordInputProps.errorMessage}
              // onChange={handleChange}
              // onBlur={handleBlur}
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
