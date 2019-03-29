import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { Link } from 'react-router-dom';
import styled, { keyframes } from 'styled-components';
import { observer, inject } from 'mobx-react';
import { UserStore } from '../stores/User';
import validator from 'validator';
import Recaptcha from 'react-recaptcha';
import media from '../components/theme/media';
import logo from '../components/assets/iecse_logo.png';
import hawk from '../components/assets/hawk_logo.png';
import pcb from '../components/assets/pcbdesign.png';
import pcb1 from '../components/assets/pcbdesign1.png';

import { Button, TextField } from '../components';
interface IState {
	count: number;
	cursor: number;
	cursor1: number;
	cursor2: number;
	cursor3: number;
	display: string;
	fullText: string;
	display1: string;
	display2: string;
	display3: string;
	interval: NodeJS.Timeout | null;
	interval1: NodeJS.Timeout | null;
	interval2: NodeJS.Timeout | null;
	interval3: NodeJS.Timeout | null;
	introq: number;
	sideq: number;
	edenq: number;
	firstq: number;
}

import Installation09Icon from '../components/assets/installation_09.svg';
import PancheaIcon from '../components/assets/panchea.svg';
import CityOfDarwinsIcon from '../components/assets/city_of_darwins.svg';
import LakesideRuinsIcon from '../components/assets/lakeside_ruins.svg';
import AshValleyIcon from '../components/assets/ash_valley.svg';
import EdenIcon from '../components/assets/eden.svg';

const size = {
	mobileS: '320px',
	tablet: '768px',
	laptop: '730px',
	laptopL: '862px',
	desktop: '1000px'
};
export const device = {
	mobileS: `(min-width: ${size.mobileS})`,
	tablet: `(min-width: ${size.tablet})`,
	laptop: `(min-width: ${size.laptop})`,
	laptopL: `(min-width: ${size.laptopL})`,
	desktop: `(min-width: ${size.desktop})`
};

const intro1 = [];

type Props = {
	className: string;
	UserStore: UserStore;
};

@inject('UserStore')
@observer
class TutPage extends Component<Props, IState> {
	public state = {
		count: 0,
		cursor: 0,
		cursor1: 0,
		cursor2: 0,
		cursor3: 0,
		display: '',
		display1: '',
		display2: '',
		display3: '',
		fullText: `Welcome to HawkEye, an online scavenger hunt based on pop culture, trivia, sports, history and just about anything else under the sun... There are four main regions - Intallation 09, Panchea, City of Darwins and Lakeside Ruins. Your quest begins in Installation09.`,
		interval: null,
		interval1: null,
		interval2: null,
		interval3: null,
		sideq: 0,
		edenq: 0,
		introq: 0,
		firstq: 1
	};

	public componentDidMount() {
		const interval = setInterval(() => {
			this.setState({
				cursor: this.state.cursor + 1,
				display: this.state.display + this.state.fullText[this.state.cursor]
			});
			if (this.state.cursor >= this.state.fullText.length) {
				clearInterval(interval);
			}
		}, 50);
		this.setState({ interval });
	}

	/************ 2nd  ****************/
	dets = () => {
		// this.setState({
		// 	cursor: 0,
		// 	sideq: 1,
		// 	display: ''
		// });
		const interval1 = setInterval(this.showww1, 50);
		this.setState({
			count: 1,
			firstq: 0,
			introq: 1,
			interval1: interval1
		});
	};

	showww1 = () => {
		let x = `When you solve all four questions in a region, a new region is unlocked for you. Once you complete all the questions in the four main regions you proceed to EDEN. Your objective is to solve the most number of questions in a fair manner...`;

		this.setState({
			cursor1: this.state.cursor1 + 1,
			display1: this.state.display1 + x[this.state.cursor1]
		});
		if (this.state.cursor1 >= x.length) {
			clearInterval(this.state.interval1!);
		}
	};

	/************ 3rd  ****************/
	sideq = () => {
		this.setState({
			count: 2,
			cursor: 0,
			introq: 0,
			sideq: 1,
			display: ''
		});
		const interval2 = setInterval(this.showww, 50);
		this.setState({
			interval2: interval2
		});
	};

	showww = () => {
		let x = `When you are stuck at a particular region, you can visit the Hermit of Ash Valley. For every question you solve you get a scroll and once you have 3 scrolls the Hermit unlocks a new region for you.`;
		this.setState({
			cursor2: this.state.cursor2 + 1,
			display2: this.state.display2 + x[this.state.cursor2]
		});
		if (this.state.cursor2 >= x.length) {
			clearInterval(this.state.interval2!);
		}
	};

	/************ 4th  ****************/
	eden = () => {
		this.setState({
			cursor: 3,
			introq: 0,
			sideq: 0,
			edenq: 1,
			display: ''
		});
		const interval3 = setInterval(this.showww2, 50);
		this.setState({
			interval3: interval3
		});
	};

	showww2 = () => {
		let x =
			'EDEN, the final region, will not be unlocked until you complete the four main regions...';
		this.setState({
			cursor3: this.state.cursor3 + 1,
			display3: this.state.display3 + x[this.state.cursor3]
		});
		if (this.state.cursor3 >= x.length) {
			console.log(this.state.count);
			clearInterval(this.state.interval3!);
		}
	};

	render() {
		return (
			<div className={this.props.className}>
				<div
					className={this.state.firstq || this.state.introq ? 'show' : 'dont'}
				>
					<img src={Installation09Icon} alt="" id="imgsiz1" />
					<img src={PancheaIcon} alt="" id="imgsiz2" />
					<img src={CityOfDarwinsIcon} alt="" id="imgsiz3" />
					<img src={LakesideRuinsIcon} alt="" id="imgsiz4" />
					<img src={AshValleyIcon} alt="" id="imgsiz5" />
					<img src={EdenIcon} alt="" id="imgsiz6" />
					<img src={hawk} alt="" id="hawklogo" />
				</div>
				<div className={this.state.sideq ? 'show' : 'dont'}>
					<img src={AshValleyIcon} alt="" id="imgsiz5new" />
				</div>
				<div className={this.state.edenq ? 'show' : 'dont'}>
					<img src={EdenIcon} alt="" id="imgsiz6new" />
				</div>

				<div id="textcontent">
					<p className={this.state.firstq ? 'show' : 'dont'}>
						{this.state.display}
					</p>
					<p className={this.state.introq ? 'show' : 'dont'}>
						{this.state.display1}
					</p>
					<p className={this.state.sideq ? 'show' : 'dont'}>
						{this.state.display2}
					</p>
					<p className={this.state.edenq ? 'show' : 'dont'}>
						{this.state.display3}
					</p>
					<button
						onClick={
							this.state.count == 0
								? this.dets
								: this.state.count == 1
								? this.sideq
								: this.eden
						}
						id={this.state.cursor3 > 0 ? 'block' : 'unblock'}
					>
						Next
					</button>
					<Link to="/dashboard">
						<button id={this.state.cursor3 > 0 ? 'unblock' : 'block'}>
							Play
						</button>
					</Link>
				</div>
			</div>
		);
	}
}

const appear = keyframes`
0%{
	opacity:0.4;
	width:20%
}
75%{
	opacity:1;
	width:25%

}
100%{
	opacity:0;
}
`;

const appear1 = keyframes`
0%{
	opacity:0.4;
	width:60%
}
75%{
	opacity:1;
	width:70%

}
100%{
	opacity:0;
}
`;

const hawkcolor = keyframes`
0%{
	opacity:0.05;
}
100%{
	opacity:1;
	filter: none;
  -webkit-filter: grayscale(0%);
}
`;

const updown = keyframes`
0% {
		transform: translate(-50%, -65%);
	}
	20% {
		transform: translate(-50%, -68%);
	}
	60% {
		transform: translate(-50%, -65%);
	}
	80% {
		transform: translate(-50%, -62%);
	}
	100% {
		transform: translate(-50%, -65%);
	}
`;

const updown1 = keyframes`
0% {
		transform: translate(-50%, -85%);
	}
	20% {
		transform: translate(-50%, -89%);
	}
	60% {
		transform: translate(-50%, -85%);
	}
	80% {
		transform: translate(-50%, -81%);
	}
	100% {
		transform: translate(-50%, -85%);
	}
`;

export default styled(TutPage)`
	.show {
		display: block;
	}
	.dont {
		display: none;
	}

	#block {
		display: none !important;
	}
	#unblock {
		display: block !important;
	}
	@media ${device.mobileS} {
		max-width: 768px;
		#imgsiz1 {
			opacity: 0;
			width: 70%;
			position: absolute;
			left: 50%;
			top: 50%;
			transform: translate(-50%, -85%);
			animation: ${appear1} 2s 1 0.1s ease-in-out;
		}
		#imgsiz2 {
			opacity: 0;
			width: 70%;
			position: absolute;
			left: 50%;
			top: 50%;
			transform: translate(-50%, -85%);
			animation: ${appear1} 2s 1 2.5s ease-in-out;
		}
		#imgsiz3 {
			opacity: 0;
			width: 70%;
			position: absolute;
			left: 50%;
			top: 50%;
			transform: translate(-50%, -85%);
			animation: ${appear1} 2s 1 5s ease-in-out;
		}
		#imgsiz4 {
			opacity: 0;
			width: 70%;
			position: absolute;
			left: 50%;
			top: 50%;
			transform: translate(-50%, -85%);
			animation: ${appear1} 2s 1 7.5s ease-in-out;
		}
		#imgsiz5 {
			opacity: 0;
			width: 70%;
			position: absolute;
			left: 50%;
			top: 50%;
			transform: translate(-50%, -85%);
			animation: ${appear1} 2s 1 10s ease-in-out;
		}
		#imgsiz5new {
			width: 80%;
			position: absolute;
			left: 50%;
			top: 50%;
			transform: translate(-50%, -85%);
			animation: ${updown1} 3s infinite linear;
		}

		#imgsiz6 {
			opacity: 0;
			width: 70%;
			position: absolute;
			left: 50%;
			top: 50%;
			transform: translate(-50%, -85%);
			animation: ${appear1} 2s 1 12.5s ease-in-out;
		}

		#imgsiz6new {
			width: 80%;
			position: absolute;
			left: 50%;
			top: 50%;
			transform: translate(-50%, -85%);
			animation: ${updown1} 3s infinite linear;
		}

		#hawklogo {
			opacity: 0.05;
			z-index: -10;
			filter: gray; /* IE6-9 */
			-webkit-filter: grayscale(100%);
			position: absolute;
			width: 90%;
			left: 50%;
			top: 50%;
			transform: translate(-50%, -80%);
			animation: ${hawkcolor} 1.5s 1 15s ease-in forwards;
		}

		#textcontent {
			color: #ffd627;
			border-radius: 10px;
			width: 70%;
			height: 40%;
			background-color: rgba(0, 0, 0, 0.4);
			position: absolute;
			left: 50%;
			bottom: 5%;
			transform: translate(-50%, 0%);
		}
		p {
			text-align: center;
			margin-left: auto;
			margin-right: auto;
			line-height: 20px;
			font-size: 0.8em;
			letter-spacing: 1.7px;
			padding: 0px 20px 0px 20px;
		}
		button {
			box-shadow: none;
			outline: none;
			-webkit-appearance: none;
			-moz-appearance: none;
			appearance: none;
			position: absolute;
			left: 50%;
			bottom: 3%;
			transform: translate(-50%, 0%);
			letter-spacing: 1.25px;
			color: #1c1c1c;
			font-weight: 500;
			background: #ffd627;
			width: 35%;
			height: 35px;
			padding: 10px;
			padding-top: 7px;
			border: none;
			border-radius: 20px;
			margin-top: 10px;
			margin-bottom: 10px;
			transition: 0.5s;
		}
		button:hover {
			color: #fff;
			background: #ff6600;
			width: 40%;
			cursor: pointer;
		}
	}

	@media ${device.tablet} {
		max-width: 1000px;
		#imgsiz1 {
			opacity: 0;
			width: 28%;
			position: absolute;
			left: 50%;
			top: 50%;
			transform: translate(-50%, -65%);
			animation: ${appear} 2s 1 0.1s ease-in-out;
		}
		#imgsiz2 {
			opacity: 0;
			width: 28%;
			position: absolute;
			left: 50%;
			top: 50%;
			transform: translate(-50%, -65%);
			animation: ${appear} 2s 1 2.5s ease-in-out;
		}
		#imgsiz3 {
			opacity: 0;
			width: 28%;
			position: absolute;
			left: 50%;
			top: 50%;
			transform: translate(-50%, -65%);
			animation: ${appear} 2s 1 5s ease-in-out;
		}
		#imgsiz4 {
			opacity: 0;
			width: 28%;
			position: absolute;
			left: 50%;
			top: 50%;
			transform: translate(-50%, -65%);
			animation: ${appear} 2s 1 7.5s ease-in-out;
		}
		#imgsiz5 {
			opacity: 0;
			width: 28%;
			position: absolute;
			left: 50%;
			top: 50%;
			transform: translate(-50%, -65%);
			animation: ${appear} 2s 1 10s ease-in-out;
		}
		#imgsiz5new {
			width: 35%;
			position: absolute;
			left: 50%;
			top: 50%;
			transform: translate(-50%, -65%);
			animation: ${updown} 3s infinite linear;
		}

		#imgsiz6 {
			opacity: 0;
			width: 28%;
			position: absolute;
			left: 50%;
			top: 50%;
			transform: translate(-50%, -65%);
			animation: ${appear} 2s 1 12.5s ease-in-out;
		}

		#imgsiz6new {
			width: 35%;
			position: absolute;
			left: 50%;
			top: 50%;
			transform: translate(-50%, -65%);
			animation: ${updown} 3s infinite linear;
		}

		#hawklogo {
			opacity: 0.05;
			z-index: -10;
			filter: gray; /* IE6-9 */
			-webkit-filter: grayscale(100%);
			position: absolute;
			width: 45%;
			left: 50%;
			top: 50%;
			transform: translate(-50%, -68%);
			animation: ${hawkcolor} 1.5s 1 15s ease-in forwards;
		}

		#textcontent {
			color: #ffd627;
			border-radius: 10px;
			width: 60%;
			height: 34%;
			background-color: rgba(0, 0, 0, 0.4);
			position: absolute;
			left: 50%;
			bottom: 3%;
			transform: translate(-50%, 0%);
		}
		p {
			text-align: center;
			margin-left: auto;
			margin-right: auto;
			line-height: 27px;
			font-size: 1em;
			letter-spacing: 1.7px;
			padding: 0px 20px 0px 20px;
		}
		button {
			box-shadow: none;
			outline: none;
			-webkit-appearance: none;
			-moz-appearance: none;
			appearance: none;
			position: absolute;
			left: 50%;
			bottom: 3%;
			transform: translate(-50%, 0%);
			letter-spacing: 1.25px;
			color: #1c1c1c;
			font-weight: 500;
			background: #ffd627;
			width: 15%;
			height: 30px;
			padding-top: 7px;
			border: none;
			border-radius: 20px;
			margin-top: 10px;
			margin-bottom: 5px;
			transition: 0.5s;
		}
		button:hover {
			color: #fff;
			background: #ff0000;
			width: 20%;
			cursor: pointer;
		}
	}

	@media ${device.desktop} {
		#imgsiz1 {
			opacity: 0;
			width: 25%;
			position: absolute;
			left: 50%;
			top: 50%;
			transform: translate(-50%, -65%);
			animation: ${appear} 2s 1 0.1s ease-in-out;
		}
		#imgsiz2 {
			opacity: 0;
			width: 25%;
			position: absolute;
			left: 50%;
			top: 50%;
			transform: translate(-50%, -65%);
			animation: ${appear} 2s 1 2.5s ease-in-out;
		}
		#imgsiz3 {
			opacity: 0;
			width: 25%;
			position: absolute;
			left: 50%;
			top: 50%;
			transform: translate(-50%, -65%);
			animation: ${appear} 2s 1 5s ease-in-out;
		}
		#imgsiz4 {
			opacity: 0;
			width: 25%;
			position: absolute;
			left: 50%;
			top: 50%;
			transform: translate(-50%, -65%);
			animation: ${appear} 2s 1 7.5s ease-in-out;
		}
		#imgsiz5 {
			opacity: 0;
			width: 25%;
			position: absolute;
			left: 50%;
			top: 50%;
			transform: translate(-50%, -65%);
			animation: ${appear} 2s 1 10s ease-in-out;
		}
		#imgsiz5new {
			width: 30%;
			position: absolute;
			left: 50%;
			top: 50%;
			transform: translate(-50%, -65%);
			animation: ${updown} 3s infinite linear;
		}

		#imgsiz6 {
			opacity: 0;
			width: 25%;
			position: absolute;
			left: 50%;
			top: 50%;
			transform: translate(-50%, -65%);
			animation: ${appear} 2s 1 12.5s ease-in-out;
		}

		#imgsiz6new {
			width: 30%;
			position: absolute;
			left: 50%;
			top: 50%;
			transform: translate(-50%, -65%);
			animation: ${updown} 3s infinite linear;
		}

		#hawklogo {
			opacity: 0.05;
			z-index: -10;
			filter: gray; /* IE6-9 */
			-webkit-filter: grayscale(100%);
			position: absolute;
			width: 40%;
			left: 50%;
			top: 50%;
			transform: translate(-50%, -68%);
			animation: ${hawkcolor} 1.5s 1 15s ease-in forwards;
		}

		#textcontent {
			color: #ffd627;
			border-radius: 10px;
			width: 60%;
			height: 30%;
			background-color: rgba(0, 0, 0, 0.4);
			position: absolute;
			left: 50%;
			bottom: 3%;
			transform: translate(-50%, 0%);
		}
		p {
			margin-left: auto;
			margin-right: auto;
			line-height: 34px;
			font-size: 1.2em;
			letter-spacing: 1.7px;
			padding: 0px 20px 0px 20px;
		}
		button {
			box-shadow: none;
			outline: none;
			-webkit-appearance: none;
			-moz-appearance: none;
			appearance: none;
			position: absolute;
			left: 50%;
			bottom: 3%;
			transform: translate(-50%, 0%);
			letter-spacing: 1.25px;
			color: #1c1c1c;
			font-weight: 500;
			background: #ffd627;
			width: 15%;
			height: 35px;
			padding: 10px;
			padding-top: 7px;
			border: none;
			border-radius: 20px;
			margin-top: 10px;
			margin-bottom: 10px;
			transition: 0.5s;
		}
		button:hover {
			color: #fff;
			background: #ff0000;
			width: 20%;
			cursor: pointer;
		}
	}
`;
