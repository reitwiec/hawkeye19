import React, { Component } from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import { Link } from 'react-router-dom';

const regions = [
	'Installation 09',
	'Panchaea',
	'City of Darwins',
	'Lakeside Ruins',
	'Ash Valley'
];

class Dashboard extends Component {
	render() {
		const { className } = this.props;
		return (
			<div className={className}>
				<h1>Dashboard</h1>
				<div className="regions-container">
					{regions.map((region, i) => (
						<span key={i}>
							<Link
								className="region-link"
								to={{ pathname: '/question', state: { region } }}
							>
								{region}
							</Link>
						</span>
					))}
				</div>
			</div>
		);
	}
}

Dashboard.propTypes = {
	className: PropTypes.string
};

export default styled(Dashboard)`
	.regions-container {
		display: flex;
		flex-flow: column nowrap;
	}

	.region-link {
		color: black;
		text-decoration: none;
	}
`;
