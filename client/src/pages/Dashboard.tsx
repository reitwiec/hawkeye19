import React, { Component } from 'react';
import styled from 'styled-components';
import { Link } from 'react-router-dom';
import { inject, observer } from 'mobx-react';

import { UserStore } from '../stores/User';

const regions = [
	'Installation 09',
	'Panchaea',
	'City of Darwins',
	'Lakeside Ruins',
	'Ash Valley'
];

type Props = {
	className: string;
	UserStore: UserStore;
};

@inject('UserStore')
@observer
class Dashboard extends Component<Props> {
	render() {
		const { className } = this.props;
		return (
			<div className={className}>
				<h1>Dashboard</h1>
				<h5>Logged in as {this.props.UserStore.username}</h5>
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

export default styled(Dashboard)`
	.regions-container,
	.attempts-container {
		display: flex;
		flex-flow: column nowrap;
	}

	.attempts-container {
		margin-top: 1em;
	}

	.region-link {
		color: black;
		text-decoration: none;
	}
`;
