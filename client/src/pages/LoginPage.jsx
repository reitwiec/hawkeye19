import React, { useCallback, useState } from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';

import { Button, TextField } from '../components';

const LoginPage = ({ className }) => {
	const [formData, setformData] = useState({
		username: '',
		password: ''
	});

	const onChange = useCallback((name, value) => {
		setformData(Object.assign(formData, { [name]: value }));
	}, []);

	const onSubmit = useCallback(() => {
		// TODO: Make API call
	}, []);

	return (
		<div className={className}>
			<h1>Login</h1>
			<TextField name="username" placeholder="Username" onChange={onChange} />
			<TextField
				name="password"
				type="password"
				placeholder="Password"
				onChange={onChange}
			/>
			<Button onClick={onSubmit}>Login</Button>
		</div>
	);
};

LoginPage.propTypes = {
	className: PropTypes.string
};

export default styled(LoginPage)`
	display: flex;
	flex-flow: column nowrap;
	width: min-content;
`;
