import React from 'react';
import { Redirect, Route } from 'react-router-dom';

const redirect = '/dashboard';

const PrivateRoute = ({ component: Component, auth, ...restProps }) => (
	<Route
		{...restProps}
		render={props =>
			!auth() ? (
				<Component {...props} />
			) : (
				<Redirect
					to={{ pathname: redirect, state: { from: props.location } }}
				/>
			)
		}
	/>
);

export default PrivateRoute;
