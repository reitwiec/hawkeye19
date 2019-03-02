import React, { useCallback, useState } from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';

import { Link } from 'react-router-dom';
import { withCookies, Cookies } from 'react-cookie';

import { Button, TextField } from '../components';

const LoginPage = ({ className, cookies }) => {
	const [formData, setformData] = useState({
		username: '',
		password: ''
	});
	const [loggedIn, setLoggedIn] = useState(false);

	const onChange = useCallback((name, value) => {
		setformData(Object.assign(formData, { [name]: value }));
	}, []);

	const login = () => {
		fetch('/api/login', {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify(formData)
		})
			.then(res => res.json())
			.then(json => setLoggedIn(json.success));
	};

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
			<Button onClick={login}>Login</Button>
			<Link to="/register">Register</Link>
		</div>
	);
};

LoginPage.propTypes = {
	className: PropTypes.string,
	cookies: PropTypes.instanceOf(Cookies)
};

export default styled(withCookies(LoginPage))`
	display: flex;
	flex-flow: column nowrap;
	width: min-content;
`;
