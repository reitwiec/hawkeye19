import React from 'react';
import { Redirect, Route } from 'react-router-dom';

import Logout from './Logout';

const redirect = '/';

const PrivateRoute = ({ component: Component, auth, ...restProps }) => (
	<Route
		{...restProps}
		render={props =>
			auth() ? (
				<>
					<Logout />
					<Component {...props} />
				</>
			) : (
				<Redirect
					to={{ pathname: redirect, state: { from: props.location } }}
				/>
			)
		}
	/>
);

export default PrivateRoute;
