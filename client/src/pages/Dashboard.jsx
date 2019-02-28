import React, { Component } from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';

class Dashboard extends Component {
	render() {
		const { className } = this.props;
		return (
			<div className={className}>
				<h1>Dashboard</h1>
			</div>
		);
	}
}

Dashboard.propTypes = {
	className: PropTypes.string
};

export default styled(Dashboard)``;
