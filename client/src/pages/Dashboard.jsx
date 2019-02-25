import React, { Component } from 'react';
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

export default styled(Dashboard)``;
