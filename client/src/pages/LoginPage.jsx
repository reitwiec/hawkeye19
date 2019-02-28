import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';

const LoginPage = ({ className }) => (
	<div className={className}>
		<h1>Login</h1>
	</div>
);

LoginPage.propTypes = {
	className: PropTypes.string
};

export default styled(LoginPage)``;
