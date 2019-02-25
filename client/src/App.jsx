import React, { Component } from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';

import { LoginPage, Dashboard } from './pages';
import { AdminRouter } from './routers';

class App extends Component {
	render() {
		return (
			<div>
				<Router>
					<Switch>
						<Route exact path="/" render={() => <LoginPage />} />
						<Route exact path="/dashboard" render={() => <Dashboard />} />
						<Route path="/___admin" component={AdminRouter} />
					</Switch>
				</Router>
			</div>
		);
	}
}

export default App;
