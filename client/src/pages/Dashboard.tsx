import React, { Component } from 'react';
import styled, { css } from 'styled-components';
import { Link } from 'react-router-dom';
import { inject, observer } from 'mobx-react';
import { UserStore } from '../stores/User';
import { RegionCard } from '../components';
import media from '../components/theme/media';
import colors from '../colors';
import { onCircle } from '../mixins';
import { RegionKey } from '../utils';

import Installation09Icon from '../components/assets/installation_09.svg';
import PancheaIcon from '../components/assets/panchea.svg';
import CityOfDarwinsIcon from '../components/assets/city_of_darwins.svg';
import LakesideRuinsIcon from '../components/assets/lakeside_ruins.svg';
import AshValleyIcon from '../components/assets/ash_valley.svg';
import EdenIcon from '../components/assets/eden.svg';

const size = {
	mobileS: '320px',
	mobileL: '500px',
	tablet: '768px',
	laptop: '730px',
	laptopL: '862px',
	desktop: '1000px'
};
export const device = {
	mobileS: `(min-width: ${size.mobileS})`,
	mobileL: `(min-width: ${size.mobileL})`,
	tablet: `(min-width: ${size.tablet})`,
	laptop: `(min-width: ${size.laptop})`,
	laptopL: `(min-width: ${size.laptopL})`,
	desktop: `(min-width: ${size.desktop})`
};

interface IDashBoardProps {
	className: string;
	UserStore: UserStore;
}

@inject('UserStore')
@observer
class Dashboard extends Component<IDashBoardProps> {
	state = {
		regions: [
			{
				key: RegionKey.region1,
				name: 'Installation 09',
				icon: Installation09Icon,
				locked: false,
				level: 1
			},
			{
				key: RegionKey.region2,
				name: 'Panchea',
				icon: PancheaIcon,
				locked: true,
				level: 0
			},
			{
				key: RegionKey.region3,
				name: 'City of Darwins',
				icon: CityOfDarwinsIcon,
				locked: true,
				level: 0
			},
			{
				key: RegionKey.region4,
				name: 'Lakeside Ruins',
				icon: LakesideRuinsIcon,
				locked: true,
				level: 0
			},
			{
				key: RegionKey.region5,
				name: 'Eden',
				icon: EdenIcon,
				locked: true,
				level: 0
			}
		]
	};

	componentDidMount() {
		this.getUser();
		setTimeout(() => this.setState({
				regions: [
					...this.state.regions.map(region => ({
						...region,
						locked: this.props.UserStore[region.key] === 0 ? true : false,
						level: this.props.UserStore[region.key]
					}))
				]
			},
			() =>
				this.setState({
					regions: this.state.regions.filter(region => !region.locked)
				})
		), 500);
	}

	getUser = () => {
		fetch(`/api/getUser`)
		.then(res => res.json())
		.then(json => {
			const userFields = {
				username: json.data.username,
				email: json.data.email,
				region0: json.data.region0,
				region1: json.data.region1,
				region2: json.data.region2,
				region3: json.data.region3,
				region4: json.data.region4,
				region5: json.data.region5,
			}
			this.props.UserStore.setCurrentUser(userFields);
		});
	};

	render() {
		const { className } = this.props;
		return (
			<div className={className}>
				<h1>Dashboard</h1>
				<h5>Logged in as {this.props.UserStore.username}</h5>
				<div className="regions-container">
					{this.state.regions.map((region, i) => (
						<RegionCard key={i} regionIndex={i + 1} {...region} />
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
		${onCircle(5, '300px', '90px', 270)}
	}

	@media ${device.mobileS} {
		max-width: 768px;
		.regions-container {
			display: flex;
			flex-flow: column nowrap;
			border: none;

			> .region-card {
				transform: unset;
				position: unset;
				width: 250px;
				height: 250px;
				margin: auto;
				margin-bottom: 4em;

				.name {
					margin-top: 10px;
				}
			}
		}
	}

	@media ${device.tablet} {
	}
`;
