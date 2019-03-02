import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';

const QuestionPage = ({ className, ...props }) => {
	const { region } = props.location.state || { region: 'Installation 09' };
	return (
		<div className={className}>
			<h1>{region}</h1>
		</div>
	);
};

QuestionPage.propTypes = {
	className: PropTypes.string
};

export default styled(QuestionPage)``;
