import React, { Component } from 'react';
import styled, { keyframes } from 'styled-components';
import map from '../components/assets/hawkbg.png';
import logo from '../components/assets/hawk_logo.png';
import { Control } from '../components';
import PropTypes from 'prop-types';
import { Link } from 'react-router-dom';
import AshValleyIcon from '../components/assets/ash_valley.svg';
interface ISideQuestProps {
	className: string;
}

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
interface IQuestionPageState {
	points: number;
	question: string;
	questionID: number | null;
	answer: string;
	hawkMessage: string;
	// attempts: string[];
	// hints: string[];
	// barOpen: boolean;
	// snackbarMessage: string;
}
const hawkResponses = {
	1: 'Hawk approves',
	2: "Hawk thinks you're close",
	3: 'Hawk disapproves'
};

class Sidequest extends Component<ISideQuestProps, IQuestionPageState> {
	state = {
		points: 0,
		question: '',
		questionID: null,
		answer: '',
		hawkMessage: ''

		// barOpen: false,
		// snackbarMessage: ''
	};

	getQuestion = (after = () => {}) => {
		fetch(`/api/getSidequestQuestion`)
			.then(res => res.json())
			.then(json => {
				this.setState(
					{
						question: json.data.question,
						questionID: json.data.questionID
					},
					() => {
						after();
					}
				);
			})
			.catch(() => {});
	};

	getUser = (after = () => {}) => {
		fetch(`/api/getUser`)
			.then(res => res.json())
			.then(json => {
				this.setState(
					{
						points: json.data.sideQuestPoints
					},
					() => {
						after();
					}
				);
			})
			.catch(() => {});
	};
	render() {
		return (
			<div className={this.props.className}>
				<div id="points">{`Points : ${this.state.points}`}</div>
				<div className="container1">
					<div id="hints" className="hintnew">
						<div className="tab">
							<button className="tablinks" id="active">
								Hints
							</button>
						</div>
						<div id="hints" className="available">
							{/* {this.state.hints.map(function(item, i) {
								return <p key={i}>{item}</p>;
							})} */}
						</div>
					</div>

					<div id="questionbox">
						<div id="level">{`Ash Valley`}</div>
						<img src={AshValleyIcon} alt="" id="ash" />
						<div id="question">{this.state.question}</div>
						<span id="status"> {this.state.hawkMessage} </span>
						<div id="answerbox">
							<input
								type="text"
								id="answer"
								placeholder="Enter answer here..."
								// value={this.state.answer}
								// onChange={this.onEditAnswer}
							/>
							<button>Submit</button>
						</div>
					</div>

					<div id="attempts" className="hintnew">
						<div className="tab">
							<button className="tablinks" id="active">
								Hints
							</button>
						</div>
						<div id="hints" className="available">
							{/* {this.state.hints.map(function(item, i) {
								return <p key={i}>{item}</p>;
							})} */}
						</div>
					</div>
				</div>
				<Control />
			</div>
		);
	}
}

export default styled(Sidequest)`
	@media ${device.mobileS} {
		#points {
			left: 50%;
			top: 15% !important;
			transform: translate(-50%, 0);
			position: absolute;
			font-size: 1.2em;
			font-weight: 500;
			letter-spacing: 1.5px;
			top: 20px;
			color: white;
		}
		#ash {
			width: 80%;
			opacity: 0.1;
			/* z-index: -10; */
			filter: gray; /* IE6-9 */
			-webkit-filter: grayscale(100%);
		}
		#status {
			position: absolute;
			bottom: 70px;
			left: 50%;
			color: #ffd627;
			transform: translate(-50%, 0%);
		}
		.container1 {
			height: 100%;
			display: flex;
			align-items: center;
			justify-content: center;
		}
		button {
			background: #181818;
			color: white;
			border: 0px;
			border-radius: 4px;
			box-shadow: none;
			outline: none;
			-webkit-appearance: none;
			-moz-appearance: none;
			appearance: none;
			height: 2.7em;
			font-size: 0.8em;
			letter-spacing: 1px;
			position: absolute;
			right: 4%;
			top: 50%;
			transform: translate(0%, -60%);
		}
		button:hover {
			cursor: pointer;
		}
		#level {
			font-size: 1.2em;
			font-weight: 700;
			color: #fff;
			margin-top: 10px;
		}
		#questionbox {
			left: 50%;
			top: 50%;
			transform: translateX(-50%);
			transform: translateY(40%);
			overflow: hidden;
			border-radius: 10px;
			background: #242121;
			filter: drop-shadow(0px 15px 15px #000);
			width: 75%;
			height: 60vh;
			text-align: center;
			z-index: 100;
			#question {
				font-size: 2vw;
				position: absolute;
				left: 50%;
				top: 10%;
				transform: translate(-50%, 0%);
				width: 80%;
				letter-spacing: 1px;
				display: block;
				text-indent: 10px;
				font-weight: 500;
				color: #fff;
			}

			#answerbox {
				position: absolute;
				bottom: 0;
				height: 60px;
				background: #ffd627;
				width: 100%;
			}
			#answer {
				position: absolute;
				left: 40%;
				top: 50%;
				transform: translate(-50%, -60%);
				width: 65%;
				height: 2.7em;
				font-size: 0.8em;
				letter-spacing: 1px;
				display: block;
				text-indent: 10px;
				font-weight: 500;
				border: 0px;
				border-radius: 4px;
				box-shadow: none;
				outline: none;
				-webkit-appearance: none;
				-moz-appearance: none;
				appearance: none;
				background: rgba(255, 255, 255, 0.3);
				text-transform: capitalize;
			}
		}
		${Control} #sideq {
			width: 100%;
		}
		${Control} #hawklogo {
			width: 15%;
		}
		${Control} .fa-question {
			font-size: 10vw;
			right: 19%;
		}
		${Control} .fa-chess-rook {
			font-size: 10vw;
			left: 19%;
		}
	}
	@media ${device.mobileL} {
	}
	@media ${device.tablet} {
	}
	@media ${device.desktop} {
		#points {
			left: 50%;
			top: 15% !important;
			transform: translate(-50%, 0);
			position: absolute;
			font-size: 1.2em;
			font-weight: 500;
			letter-spacing: 1.5px;
			top: 20px;
			color: white;
		}
		#ash {
			width: 60%;
			opacity: 0.1;
			/* z-index: -10; */
			filter: gray; /* IE6-9 */
			-webkit-filter: grayscale(100%);
		}
		#status {
			position: absolute;
			bottom: 70px;
			left: 50%;
			color: #ffd627;
			transform: translate(-50%, 0%);
		}
		.container1 {
			height: 100%;
			display: flex;
			align-items: center;
			justify-content: center;
		}
		button {
			background: #181818;
			color: white;
			border: 0px;
			border-radius: 4px;
			box-shadow: none;
			outline: none;
			-webkit-appearance: none;
			-moz-appearance: none;
			appearance: none;
			height: 2.7em;
			font-size: 0.8em;
			letter-spacing: 1px;
			position: absolute;
			right: 4%;
			top: 50%;
			transform: translate(0%, -60%);
		}
		button:hover {
			cursor: pointer;
		}
		#level {
			font-size: 1.2em;
			font-weight: 700;
			color: #fff;
			margin-top: 10px;
		}
		#questionbox {
			left: 50%;
			top: 50%;
			transform: translateX(-50%);
			transform: translateY(50%);
			overflow: hidden;
			border-radius: 10px;
			background: #242121;
			filter: drop-shadow(0px 15px 15px #000);
			width: 35%;
			height: 380px;
			text-align: center;
			z-index: 100;
			#question {
				font-size: 2vw;
				position: absolute;
				left: 50%;
				top: 10%;
				transform: translate(-50%, 0%);
				width: 80%;
				letter-spacing: 1px;
				display: block;
				text-indent: 10px;
				font-weight: 500;
				color: #fff;
			}

			#answerbox {
				position: absolute;
				bottom: 0;
				height: 60px;
				background: #ffd627;
				width: 100%;
			}
			#answer {
				position: absolute;
				left: 40%;
				top: 50%;
				transform: translate(-50%, -60%);
				width: 70%;
				height: 2.7em;
				font-size: 0.8em;
				letter-spacing: 1px;
				display: block;
				text-indent: 10px;
				font-weight: 500;
				border: 0px;
				border-radius: 4px;
				box-shadow: none;
				outline: none;
				-webkit-appearance: none;
				-moz-appearance: none;
				appearance: none;
				background: rgba(255, 255, 255, 0.3);
				text-transform: capitalize;
			}
		}
		${Control} #sideq {
			width: 30%;
		}
		${Control} #hawklogo {
			width: 4%;
		}
		${Control} .fa-question {
			font-size: 3vw;
			right: 40%;
		}
		${Control} .fa-chess-rook {
			font-size: 3vw;
			left: 40%;
		}

		/* #hints {
			transform: translateY(50%);
			margin-right: 50px;
			overflow: hidden;
			border-radius: 10px;
			background: #242121;
			filter: drop-shadow(0px 15px 15px #000);
			width: 25%;
			height: 380px;
			text-align: center;
			z-index: 100;
		} */
	}
`;
