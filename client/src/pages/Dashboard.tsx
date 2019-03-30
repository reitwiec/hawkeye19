import React, { Component } from 'react';
import styled, { css } from 'styled-components';
import { Link } from 'react-router-dom';
import { inject, observer } from 'mobx-react';
import { UserStore } from '../stores/User';
import { RegionCard } from '../components';
import media from '../components/theme/media';
import colors from '../colors';
import { onCircle } from '../mixins';
import logo from '../components/assets/hawk_logo.png';
import sideq from '../components/assets/sideq.svg';
import { RegionKey } from '../utils';
import iecse from '../components/assets/iecse_logo.png';
import Installation09Icon from '../components/assets/installation_09.svg';
import PancheaIcon from '../components/assets/panchea.svg';
import CityOfDarwinsIcon from '../components/assets/city_of_darwins.svg';
import LakesideRuinsIcon from '../components/assets/lakeside_ruins.svg';
import AshValleyIcon from '../components/assets/ash_valley.svg';
import EdenIcon from '../components/assets/eden.svg';
import { Redirect } from 'react-router-dom';
import { Logout } from '../components';

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
	history: any;
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
		],
		redirect: null
	};

	componentDidMount() {
		this.getUser(this.initset);
	}

	unlock = (name, regionID) => e => {
		this.props.UserStore.setRegion(regionID);
		this.setState({
			redirect: (
				<Redirect to={{ pathname: '/question', state: { name, regionID } }} />
			)
		});
	};

	locky = () => {};

	getUser = (callback = () => {}) => {
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
					isVerified: json.data.isVerified,
					sideQuestPoints: json.data.sideQuestPoints
				};
				this.props.UserStore.setCurrentUser(userFields);
				callback();
			});
	};

	initset = () => {
		this.setState({
			regions: [
				...this.state.regions.map(region => ({
					...region,
					locked: this.props.UserStore[region.key] === 0 ? true : false,
					level: this.props.UserStore[region.key]
				}))
			]
		});
	};

	render() {
		const { className } = this.props;
		return (
			<div className={className}>
				<img src={iecse} alt="" id="iecse" />
				<h1>Dashboard</h1>
				{/* <h5>Logged in as {this.props.UserStore.username}</h5> */}
				<div className="regions-container">
					{this.state.regions.map((region, i) => (
						<div
							className="questions"
							onClick={
								this.state.regions[i].locked || this.state.regions[i].level > 4
									? this.locky
									: this.unlock(this.state.regions[i].name, i + 1)
							}
							id={
								this.state.regions[i].level > 4 ? 'completed' : 'notcompleted'
							}
						>
							<section className={this.state.regions[i].locked ? 'locked' : ''}>
								{this.state.regions[i].name}{' '}
								<img id="icons" src={this.state.regions[i].icon} alt="" />
							</section>
						</div>
						// <RegionCard key={i} regionIndex={i + 1} {...region} />
					))}
				</div>
				{this.state.redirect}
				<div id="control">
					<div id="signals">
						<Link to="/rules">
							<i className="fas fa-question" />
						</Link>
						<Logout className="logout-btn" />
						<i
							className="fas fa-chess-rook"
							onClick={this.unlock('ashvalley', 0)}
						/>
					</div>
					<img src={sideq} alt="" id="sideq" />
				</div>
			</div>
		);
	}
}

export default styled(Dashboard)`
	.logout-btn {
		position: absolute;
		bottom: 10px;
		transform: translate(-50%, 0%);
		z-index: 102;
		left: 50%;
	}
	#completed {
		user-select: none;
		background: #ffd627;
	}
	#completed > section {
		color: #131212;
	}

	@media ${device.mobileS} {
		.locked {
			user-select: none;
			filter: gray; /* IE6-9 */
			-webkit-filter: grayscale(100%);
		}
		.locked:hover {
		}
		#iecse {
			width: 10%;
			position: absolute;
			left: 2%;
			top: 2%;
		}
		#icons {
			float: right;
			width: 14%;
		}
		h1 {
			text-align: center;
			color: #ffd627;
			letter-spacing: 1.9px;
			left: 50%;
			top: 4%;
			text-transform: uppercase;
			/* transform: translate(-50%, 0%); */
		}
		.questions {
			width: 90%;
			left: 50%;
			transform: translateX(-50%);
			background: #131212;
			position: relative;
			/* color: #445076; */
			font-size: 1.3em;
			font-weight: 700;
			line-height: 2;
			padding: 12px;
			border-radius: 10px;
			/* background:#2f2e4d; */
			margin: 20px 0 20px 0;
			transition: 0.1s;
		}
		.questions > section {
			font-size: 1.2em;
			transition: 0.2s;
			color: #a7a6a6;
			letter-spacing: -0.5px;
		}
		.questions:hover {
			cursor: pointer;
			width: 93%;
			z-index: 1;
			background: #e2b700;
			// background: linear-gradient(45deg, #fd6b9a, #f77f6e);
			filter: drop-shadow(0px 5px 5px #181e30);
			color: #fff;
			section {
				font-size: 1.5em;
				color: #fff;
			}
		}
		.questions > section:hover {
			cursor: pointer;
			font-size: 1.5em;
			color: #fff;
		}
		.regions-container {
			margin-top: 5%;
		}

		#sideq {
			z-index: 20;
			position: fixed;
			bottom: 0px;
			transform: translate(-50%, 0px);
			left: 50%;
			width: 300px;
			filter: drop-shadow(rgb(0, 0, 0) 0px -5px 10px);
		}
		#control {
			position: fixed;
			bottom: 0px;
		}
		.fa-power-off {
			font-size: 1.3rem;
			background: #181818;
			left: 50%;
			color: white;
			border: 5px solid #c49905;
			padding: 9px;
			border-radius: 100px;
			bottom: 20px;
			left: 50%;
			transform: translate(-50%, 0%);
			z-index: 102;
			transition: 0.3s;
			position: fixed;
		}

		.fa-question {
			color: #242121;
			transition: 0.3s;
			font-size: 30px;
			z-index: 100;
			position: fixed;
			bottom: 0px;
			right: calc(50% - 103px);
			padding: 10px;
		}
		.fa-chess-rook {
			color: #242121;
			transition: 0.3s;
			font-size: 30px;
			z-index: 100;
			position: fixed;
			left: calc(50% - 103px);
			bottom: 0px;
			padding: 10px;
		}
	}

	/******************************************/
	@media ${device.desktop} {
		.locked {
			user-select: none;
			filter: gray; /* IE6-9 */
			-webkit-filter: grayscale(100%);
		}
		.locked:hover {
		}
		#iecse {
			width: 4%;
			position: absolute;
			left: 2%;
			top: 2%;
		}
		#icons {
			float: right;
			width: 5%;
		}
		h1 {
			text-align: center;
			color: #ffd627;
			letter-spacing: 1.9px;
			left: 50%;
			top: 4%;
			text-transform: uppercase;
			/* transform: translate(-50%, 0%); */
		}
		.questions {
			width: 90%;
			left: 50%;
			transform: translateX(-50%);
			background: #131212;
			position: relative;
			/* color: #445076; */
			font-size: 1.3em;
			font-weight: 700;
			line-height: 2;
			padding: 12px;
			border-radius: 10px;
			/* background:#2f2e4d; */
			margin: 20px 0 20px 0;
			transition: 0.1s;
		}
		.questions > section {
			font-size: 1.2em;
			transition: 0.2s;
			color: #a7a6a6;
			letter-spacing: -0.5px;
		}
		.questions:hover {
			cursor: pointer;
			width: 93%;
			z-index: 1;
			background: #e2b700;
			// background: linear-gradient(45deg, #fd6b9a, #f77f6e);
			filter: drop-shadow(0px 5px 5px #181e30);
			color: #fff;
			section {
				font-size: 1.5em;
				color: #fff;
			}
		}
		.questions > section:hover {
			cursor: pointer;
			font-size: 1.5em;
			color: #fff;
		}
		.regions-container {
			margin-top: 5%;
		}

		.fa-power-off {
			font-size: 1.9rem;
			background: #181818;
			left: 50%;
			color: white;
			border: 5px solid #c49905;
			padding: 9px;
			border-radius: 100px;
			position: fixed;
			bottom: 20px;
			left: 50%;
			transform: translate(-50%, 0%);
			z-index: 102;
			transition: 0.3s;
		}
		.fa-power-off:hover {
			cursor: pointer;
			font-size: 3.5vw;
		}
		#sideq {
			z-index: 20;
			position: fixed;
			bottom: 0px;
			transform: translate(-50%, 0);
			left: 50%;
			width: 400px;
			filter: drop-shadow(0px -5px 10px #000);
		}
		#control {
			position: fixed;
			bottom: 0px;
		}
		.fa-question {
			font-size: 38px;
			z-index: 104;
			color: #242121;
			position: fixed;
			bottom: 0;
			right: calc(50% - 110px);
			transition: 0.3s;
			/* transform: translate(-50%,0%); */
			padding: 15px;
		}
		.fa-chess-rook {
			font-size: 38px;
			z-index: 110;
			color: #242121;
			position: fixed;
			left: calc(50% - 110px);
			bottom: 0px;
			transition: 0.3s;
			padding: 15px;
		}

		.fa-chess-rook:hover,
		.fa-question:hover,
		#hawklogo:hover {
			cursor: pointer;
			font-size: 3.2vw;
		}
	}
`;
