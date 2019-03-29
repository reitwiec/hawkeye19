import React, { Component } from 'react';
import { Redirect } from 'react-router-dom';
import styled, { keyframes } from 'styled-components';

import Button from './Button';
import { UserStore } from '../stores/User';
type Props = {
	className: string;
};

class Logout extends Component<Props> {
	state = {
		loggedIn: true
	};

	logout = () => {
		fetch('/api/logout', {
			method: 'GET',
			headers: { 'Content-Type': 'application/json' }
		})
			.then(res => res.json())
			.then(json => this.setState({ loggedIn: !json.success }));
	};

	render() {
		if (this.state.loggedIn) {
			return (
				<i
					className={`fas fa-power-off ${this.props.className}`}
					onClick={this.logout}
				/>
			);
		} else {
			return <Redirect to="/" />;
		}
	}
}

export default styled(Logout)``;
