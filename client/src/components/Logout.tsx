import React, { Component } from 'react';
import { Redirect } from 'react-router-dom';

import Button from './Button';
import { UserStore } from '../stores/User';

class Logout extends Component {
	state = {
		loggedIn: true
	};

	logout = () => {
		fetch('/api/logout', {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' }
		})
			.then(res => res.json())
			.then(json => this.setState({ loggedIn: !json.success }));
	};

	render() {
		if (this.state.loggedIn) {
			return <Button onClick={this.logout}>Logout</Button>;
		} else {
			return <Redirect to="/" />;
		}
	}
}

export default Logout;
