import React, { useCallback, useState } from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';

import { Button, TextField } from '../components';

const RegisterPage = ({ className }) => {
	const [formData, setformData] = useState({
		name: '',
		username: '',
		password: '',
		confirm_password: '',
		email: '',
		tel: '',
		college: ''
	});

	const onChange = useCallback((name, value) => {
		setformData(Object.assign(formData, { [name]: value }));
	}, []);

	const onSubmit = useCallback(() => {
		// TODO: Make API call
	}, []);

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
			/>
			<TextField name="tel" placeholder="Phone Number" onChange={onChange} />
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
