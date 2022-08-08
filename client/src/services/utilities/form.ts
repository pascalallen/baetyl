import { FormikValues, FormikState } from 'formik';

type InputProps = {
  errorMessage?: string;
  isValid: boolean;
};

const getInputProps = (
  fieldName: string,
  formik: FormikState<FormikValues>,
  serverErrors: { [fieldName: string]: string }
): InputProps => {
  let error = undefined;
  if (formik.touched[fieldName] && formik.errors[fieldName]) {
    error = formik.errors[fieldName] as string;
  }

  if (serverErrors[fieldName] !== '') {
    error = serverErrors[fieldName];
  }

  return {
    errorMessage: error,
    isValid: !error
  };
};

export default Object.freeze({
  getInputProps
});
