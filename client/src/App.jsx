import React, { Component } from 'react';
import {
	BrowserRouter as Router,
	Switch
} from 'react-router-dom';
import { CookiesProvider, withCookies } from 'react-cookie';
import { Provider } from 'mobx-react';

import {
	Dashboard,
	LoginPage,
	QuestionPage,
	RegisterPage,
	SupportPage,
	TutPage,
	Rules
} from './pages';
import { PrivateRoute, PublicRoute } from './components';
import { AdminRouter } from './routers';
import * as stores from './stores';

class App extends Component {
	loggedIn = () => (this.props.cookies.get('session') ? true : false);

	render() {
		return (
			<div>
				<Provider {...stores}>
					<CookiesProvider>
						<Router>
							<Switch>
								<PublicRoute
									exact
									path="/"
									component={LoginPage}
									auth={this.loggedIn}
								/>
								<PublicRoute
									exact
									path="/register"
									component={RegisterPage}
									auth={this.loggedIn}
								/>
								<PublicRoute
									exact
									path="/tutorial"
									component={TutPage}
									auth={this.loggedIn}
								/>
								<PrivateRoute
									exact
									path="/rules"
									component={Rules}
									auth={this.loggedIn}
								/>

								<PrivateRoute
									exact
									path="/dashboard"
									component={Dashboard}
									auth={this.loggedIn}
								/>
								<PrivateRoute
									exact
									path="/question"
									component={QuestionPage}
									auth={this.loggedIn}
								/>
								<PrivateRoute
									path="/___admin"
									component={AdminRouter}
									auth={this.loggedIn}
								/>
								<PrivateRoute
									exact
									path="/support"
									component={SupportPage}
									auth={this.loggedIn}
								/>
							</Switch>
						</Router>
					</CookiesProvider>
				</Provider>
			</div>
		);
	}
}

export default withCookies(App);
