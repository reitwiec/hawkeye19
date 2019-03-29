import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';

const Button = ({ children, ...props }) => {
	return <button {...props}>{children}</button>;
};

Button.propTypes = {
	className: PropTypes.string,
	onClick: PropTypes.func.isRequired,
	disabled: PropTypes.bool
};

export default styled(Button)`
	box-shadow: none;
	outline: none;
	-webkit-appearance: none;
	-moz-appearance: none;
	appearance: none;
	:hover {
		cursor: pointer;
	}
`;
