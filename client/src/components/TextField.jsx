import React, { useState } from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';

const TextField = ({
	className,
	onChange,
	type = 'text',
	validation = null,
	validateOnChange = false,
	...props
}) => {
	const [value, setValue] = useState('');
	const [error, setError] = useState('');

	const showError = value => {
		if (validation) {
			setError(validation(value));
		}
	};

	return (
		<span className={className}>
			<input
				{...props}
				type={type}
				value={value}
				onChange={e => {
					e.preventDefault();
					setValue(e.target.value);
					if (validateOnChange && error) {
						showError(value);
					}
					onChange(props.name, e.target.value, error);
				}}
				onBlur={() => showError(value)}
			/>
			<div className={`error-label ${error ? 'error' : ''}`}>{error}</div>
		</span>
	);
};

TextField.propTypes = {
	name: PropTypes.string.isRequired,
	type: PropTypes.oneOf(['text', 'password', 'email']),
	className: PropTypes.string,
	placeholder: PropTypes.string,
	onChange: PropTypes.func.isRequired,
	validation: PropTypes.func,
	validateOnChange: PropTypes.bool
};

export default styled(TextField)`
	margin-bottom: 0.8em;
	.error-label {
		display: ${({ validation }) => (validation ? 'block' : 'none')};
		font-size: 0.8em;
		height: 1em;
	}

	.error {
		background-color: rgba(255, 0, 0, 0.5);
	}
`;
