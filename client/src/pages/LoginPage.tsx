import React, { Component } from 'react';
import styled from 'styled-components';
import { observer, inject } from 'mobx-react';
import { UserStore } from '../stores';

import { Link, Redirect } from 'react-router-dom';

import { Button, TextField } from '../components';

type Props = {
	className: string;
	UserStore: typeof UserStore;
};

@inject('UserStore')
@observer
class LoginPage extends Component<Props> {
	state = {
		username: '',
		password: '',
		loggedIn: false
	};

	onChange = (name, value) => {
		this.setState({ [name]: value });
	};

	login = () => {
		const { loggedIn, ...loginData } = this.state;
		fetch('/api/login', {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify(loginData)
		})
			.then(res => res.json())
			.then(json => {
				if (json.success) {
					this.setState({ loggedIn: true });
					this.props.UserStore.setCurrentUser(json.data);
				}
			});
	};

	render() {
		return (
			<div className={this.props.className}>
				<Link to="/register">Register</Link>
				<h1>Login</h1>
				<TextField
					name="username"
					placeholder="Username"
					onChange={this.onChange}
				/>
				<TextField
					name="password"
					type="password"
					placeholder="Password"
					onChange={this.onChange}
				/>
				<Button onClick={this.login}>Login</Button>
				{this.state.loggedIn ? <Redirect to="/dashboard" /> : null}
			</div>
		);
	}
}

export default styled(LoginPage)`
	display: flex;
	flex-flow: column nowrap;
	width: min-content;
`;
