import React, { Component } from 'react';
import { Switch } from 'react-router-dom';

import { Cookies, withCookies } from 'react-cookie';

import { QuestionPage } from '../pages/admin';
import { PrivateRoute } from '../components';

type Props = {
	cookies: Cookies;
};

class AdminRouter extends Component<Props> {
	loggedIn = () => (this.props.cookies.get('session') ? true : false);
	// loggedIn = () => true;

	render() {
		return (
			<>
				<h1>Admin</h1>
				<Switch>
					<PrivateRoute
						exact
						path="/___admin/questions"
						component={QuestionPage}
						auth={this.loggedIn}
						showLogout={false}
					/>
				</Switch>
			</>
		);
	}
}

export default withCookies(AdminRouter);
