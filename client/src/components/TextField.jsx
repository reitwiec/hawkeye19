import React, { useState } from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';

const TextField = ({ onChange, type = 'text', ...props }) => {
	const [value, setValue] = useState('');

	return (
		<input
			{...props}
			type={type}
			value={value}
			onChange={e => {
				e.preventDefault();
				setValue(e.target.value);
				onChange(props.name, e.target.value);
			}}
		/>
	);
};

TextField.propTypes = {
	name: PropTypes.string.isRequired,
	type: PropTypes.oneOf(['text', 'password', 'email']),
	className: PropTypes.string,
	placeholder: PropTypes.string,
	onChange: PropTypes.func.isRequired
};

export default styled(TextField)``;
