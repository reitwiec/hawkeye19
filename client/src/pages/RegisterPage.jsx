import React, { useCallback, useState } from "react";
import PropTypes from "prop-types";
import styled from "styled-components";
import validator from "validator";

import { Button, TextField } from "../components";

const RegisterPage = ({ className }) => {
  const [formData, setformData] = useState({
    name: "",
    username: "",
    password: "",
    confirm_password: "",
    email: "",
    tel: "",
    college: ""
  });

  const onChange = useCallback((name, value, error) => {
    setformData(Object.assign(formData, { [name]: { value, error } }));
    console.log(formData);
  }, []);

  const onSubmit = useCallback(() => {}, []);

  return (
    <div className={className}>
      <h1>Register</h1>
      <TextField name="name" placeholder="Name" onChange={onChange} />
      <TextField name="username" placeholder="Username" onChange={onChange} />
      <TextField
        name="password"
        type="password"
        placeholder="Password"
        onChange={onChange}
      />
      <TextField
        name="confirm_password"
        type="password"
        placeholder="Confirm password"
        onChange={onChange}
      />
      <TextField
        name="email"
        type="email"
        placeholder="Email"
        onChange={onChange}
        validation={v => (validator.isEmail(v) ? "" : "Invalid Email")}
        validateOnChange
      />
      <TextField
        name="tel"
        placeholder="Mobile Number"
        onChange={onChange}
        validation={v => (validator.isMobilePhone(v) ? "" : "Invalid Number")}
        validateOnChange
      />
      <TextField name="college" placeholder="College" onChange={onChange} />
      <Button onClick={onSubmit}>Register</Button>
    </div>
  );
};

RegisterPage.propTypes = {
  className: PropTypes.string
};

export default styled(RegisterPage)`
  display: flex;
  flex-flow: column nowrap;
  width: min-content;
`;
