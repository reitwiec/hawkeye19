import React, { Component } from 'react';
import styled, { css } from 'styled-components';
import { Link } from 'react-router-dom';
import { inject, observer } from 'mobx-react';
import { UserStore } from '../stores/User';
import { RegionCard } from '../components';
import media from '../components/theme/media';
import colors from '../colors';
import { onCircle } from '../mixins';

import Installation09Icon from '../components/assets/installation_09.svg';
import PancheaIcon from '../components/assets/panchea.svg';
import CityOfDarwinsIcon from '../components/assets/city_of_darwins.svg';
import LakesideRuinsIcon from '../components/assets/lakeside_ruins.svg';
import AshValleyIcon from '../components/assets/ash_valley.svg';
import EdenIcon from '../components/assets/eden.svg';

const regions = [
	{ name: 'Installation 09', icon: Installation09Icon, locked: false },
	{ name: 'Panchea', icon: PancheaIcon, locked: true },
	{ name: 'City of Darwins', icon: CityOfDarwinsIcon, locked: true },
	{ name: 'Lakeside Ruins', icon: LakesideRuinsIcon, locked: true },
	{ name: 'Ash Valley', icon: AshValleyIcon, locked: true },
	{ name: 'Eden', icon: EdenIcon, locked: true }
];

interface IDashBoardProps {
	className: string;
	UserStore: UserStore;
}

@inject('UserStore')
@observer
class Dashboard extends Component<IDashBoardProps> {
	render() {
		const { className } = this.props;
		return (
			<div className={className}>
				<h1>Dashboard</h1>
				<h5>Logged in as {this.props.UserStore.username}</h5>
				<div className="regions-container">
					{regions.map((region, i) => (
						<RegionCard key={i} {...region} />
					))}
				</div>
			</div>
		);
	}
}

export default styled(Dashboard)`
	color: ${colors.mediumGrey};
	.regions-container {
		margin: auto;
		border: 5px solid ${`${colors.primaryYellow}aa`};
		${onCircle(6, '400px', '90px', 270)}
	}

	${media.phone`
		.regions-container {
			display: flex;
			flex-flow: column nowrap;
			border: none;

			> .region-card {
				transform: unset;
				position: unset;
				width: 60vw;
				height: 60vw;
				margin-top: 4em;
				margin-left: 4em;
			}
		}
	`}
`;
