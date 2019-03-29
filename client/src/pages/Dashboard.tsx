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
		this.getUser();
		setTimeout(
			() =>
				this.setState(
					{
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
				),
			500
		);
	}
	unlock = (name, regionID) => e => {
		this.setState({
			redirect: (
				<Redirect to={{ pathname: '/questions', state: { name, regionID } }} />
			)
		});
	};

	locky = () => {
		console.log('locked');
	};

	getUser = () => {
		fetch(`/api/getUser`)
			.then(res => res.json())
			.then(json => {
				console.log(json);
				const userFields = {
					username: json.data.username,
					email: json.data.email,
					region0: json.data.region0,
					region1: json.data.region1,
					region2: json.data.region2,
					region3: json.data.region3,
					region4: json.data.region4,
					region5: json.data.region5,
					isVerified: json.data.isVerified
				};
				this.props.UserStore.setCurrentUser(userFields);
			});
	};

	render() {
		const { className } = this.props;
		return (
			<div className={className}>
				<img src={iecse} alt="" id="iecse" />
				<h1>Dashboard</h1>
				<h5>Logged in as {this.props.UserStore.username}</h5>
				<div className="regions-container">
					{this.state.regions.map((region, i) => (
						<div
							className="questions"
							onClick={
								this.state.regions[i].locked
									? this.locky
									: this.unlock(this.state.regions[i].name, i + 1)
							}
						>
							<section className={this.state.regions[i].locked ? 'locked' : ''}>
								{this.state.regions[i].name}{' '}
								<img id="icons" src={this.state.regions[i].icon} alt="" />
							</section>
							{/* <span className="details">Max Score: 15</span>
							<span className="strength">Difficulty: Easy</span> */}
						</div>
						// <RegionCard key={i} regionIndex={i + 1} {...region} />
					))}
				</div>
				{this.state.redirect}
				<div id="control">
					<div id="signals">
						<i className="fas fa-question" />
						<img src={logo} id="hawklogo" alt="" />
						<i className="fas fa-chess-rook" />
					</div>
					<img src={sideq} alt="" id="sideq" />
				</div>
			</div>
		);
	}
}

export default styled(Dashboard)`
	@media ${device.mobileS} {
		max-width: 768px;
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
		#hawklogo {
			width: 10%;
			background: #181818;
			left: 50%;
			border: 5px solid #c49905;
			padding: 2px;
			border-radius: 100px;
			position: absolute;
			bottom: 20px;
			transform: translate(-50%, 0%);
			z-index: 102;
			transition: 0.3s;
		}

		#sideq {
			z-index: 20;
			position: absolute;
			bottom: 0px;
			transform: translate(-50%, 0);
			left: 50%;
			/* 
        top: 50%;
         */
			width: 75%;
			filter: drop-shadow(0px -5px 10px #000);
		}
		#control {
			z-index: -65;
			/* background:red; */
			position: relative;
			top: 0px;
			height: 19vh;
		}
		.fa-question {
			font-size: 8vw;
			z-index: 104;
			color: #242121;
			position: absolute;
			bottom: 0.3vh;
			right: 25%;
			transition: 0.3s;
			/* transform: translate(-50%,0%); */
			padding: 15px;
		}
		.fa-chess-rook {
			font-size: 8vw;
			z-index: 110;
			color: #242121;
			position: absolute;
			left: 25%;
			bottom: 0.3vh;
			transition: 0.3s;
			/* transform: translate(-50%,0%); */
			padding: 15px;
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
		#hawklogo {
			width: 4%;
			background: #181818;
			left: 50%;
			border: 5px solid #c49905;
			padding: 2px;
			border-radius: 100px;
			position: absolute;
			bottom: 20px;
			transform: translate(-50%, 0%);
			z-index: 102;
			transition: 0.3s;
		}
		#hawklogo:hover {
			width: 4.5%;
		}
		#sideq {
			z-index: 20;
			position: absolute;
			bottom: 0px;
			transform: translate(-50%, 0);
			left: 50%;
			/* 
        top: 50%;
         */
			width: 30%;
			filter: drop-shadow(0px -5px 10px #000);
		}
		#control {
			z-index: -65;
			/* background:red; */
			position: relative;
			top: 0px;
			height: 18vh;
		}
		.fa-question {
			font-size: 3vw;
			z-index: 104;
			color: #242121;
			position: absolute;
			bottom: 0.3vh;
			right: 40%;
			transition: 0.3s;
			/* transform: translate(-50%,0%); */
			padding: 15px;
		}
		.fa-chess-rook {
			font-size: 3vw;
			z-index: 110;
			color: #242121;
			position: absolute;
			left: 40%;
			bottom: 0.3vh;
			transition: 0.3s;
			/* transform: translate(-50%,0%); */
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
