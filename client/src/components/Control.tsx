import React, { Component } from 'react';
import styled, { keyframes } from 'styled-components';
import logo from '../components/assets/hawk_logo.png';
import sideq from '../components/assets/sideq.svg';
import PropTypes from 'prop-types';
import { Link } from 'react-router-dom';
interface ISideQuestProps {
	className: string;
}

class Control extends Component<ISideQuestProps> {
	render() {
		return (
			<div className={this.props.className}>
				<div id="control">
					<div id="signals">
						<Link to="/rules">
							<i className="fas fa-question" />
						</Link>
						<Link to="/dashboard">
							<img src={logo} id="hawklogo" alt="" />
						</Link>
						<Link to="/sidequest">
							<i className="fas fa-chess-rook" />
						</Link>
					</div>
					<img src={sideq} alt="" id="sideq" />
					{/* <img src={map} alt="" id="map" /> */}
				</div>
			</div>
		);
	}
}

export default styled(Control)`
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
	#sideq {
		z-index: 20;
		position: absolute;
		bottom: 0px;
		transform: translate(-50%, 0);
		left: 50%;
		width: 30%;
		filter: drop-shadow(0px -5px 10px #000);
	}
	#hawklogo:hover {
		width: 4.5%;
	}
	#control {
		/* position: relative;
		top: 0px;
		height: 92vh; */
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
`;
