import React, { Component } from 'react';
import PropTypes from 'prop-types';
import styled, { keyframes } from 'styled-components';
import media from '../components/theme/media';
import logo from '../components/assets/hawk_logo.png';
import sideq from '../components/assets/sideq.svg';
import { inject, observer } from 'mobx-react';
import map from '../components/assets/mapbg.png';
import { Button } from '../components';
import { Link } from 'react-router-dom';
import Logout from '../components/Logout';
import { UserStore } from '../stores/User';
import { Snackbar } from '../components';

const size = {
	mobileS: '320px',
	mobileM: '375px',
	mobileL: '420px',
	tablet: '580px',
	laptop: '730px',
	laptopL: '862px',
	desktop: '1000px'
};

const hawkResponses = {
	1: 'Hawk approves',
	2: 'Hawk thinks you\'re close',
	3: 'Hawk disapproves'
};

export const device = {
	mobileS: `(min-width: ${size.mobileS})`,
	mobileM: `(min-width: ${size.mobileM})`,
	mobileL: `(min-width: ${size.mobileL})`,
	tablet: `(min-width: ${size.tablet})`,
	laptop: `(min-width: ${size.laptop})`,
	laptopL: `(min-width: ${size.laptopL})`,
	desktop: `(min-width: ${size.desktop})`,
	desktopL: `(min-width: ${size.desktop})`
};
const stats = ['Tries : 6969', 'On-par : 0', 'Leading : 1', 'Trailing : 69'];

interface IQuestionPageProps {
	className: string;
	location: { state: { name: string; regionIndex: number } };
  UserStore: UserStore;
  history: any;
}

interface IQuestionPageState {
	tryvisible: boolean;
	hintvisible: boolean;
	statsvisible: boolean;
	level: number;
	question: string;
	questionID: number | null;
	answer: string;
	attempts: string[];
	hints: string[];
	hawkMessage: string;
	barOpen: boolean;
	snackbarMessage: string;
}

@inject('UserStore')
@observer
class QuestionPage extends Component<IQuestionPageProps, IQuestionPageState> {
	state = {
		tryvisible: true,
		hintvisible: false,
		statsvisible: false,
		level: this.props.UserStore[`region${this.props.UserStore.activeRegion}`],
		question: '',
		questionID: null,
		answer: '',
		attempts: [],
		hints: [],
		hawkMessage: '',
		region: this.props.location.state.name,
		regionIndex: this.props.UserStore.activeRegion,
		barOpen: false,
		snackbarMessage: ''
	};

	componentDidMount() {
    if (!this.props.UserStore.activeRegion)
      this.props.history.push('/dashboard');
    else
      this.getQuestion(() => {
        this.getHints();
        this.getAttempts()
      });
	}

	getQuestion = (after = () => {}) => {
		// console.log(this.props.UserStore.activeRegion);
		fetch(`/api/getQuestion?region=${this.props.UserStore.activeRegion}`)
			.then(res => res.json())
			.then(json => {
				this.setState({
					question: json.data.question,
					questionID: json.data.questionID,
					level: json.data.level
				}, () => {
					after();
				});
			})
			.catch(() => {});
	};

	clearAnswer = () => {
		this.setState({ answer: '' });
	};

	getHints = () => [
		fetch(`/api/getHints?question=${this.state.questionID}`)
			.then(res => res.json())
			.then(json => {
				this.setState({ hints: json.data? json.data : [] });
			})
	];

	getAttempts = () => {
		fetch(`/api/getRecentTries?question=${this.state.questionID}`)
		.then(res => res.json())
		.then(json => {
			this.setState({ attempts: json.data? json.data : [] })
		});
	};

	checkAnswer = () => {
		if (this.props.UserStore.isVerified == 1) {
			fetch(`/api/checkAnswer`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					regionID: this.props.UserStore.activeRegion,
					answer: this.state.answer
				})
			})
				.then(res => res.json())
				.then(json => {
					this.onSubmit(json);
				});
		} else {
			this.openSnackbar('Please verify your email');
		}
	};

	onSubmit = (json) => {
		this.clearAnswer();
		this.setState({
			hawkMessage: hawkResponses[json.data]
    });
		if (json.data == 1) {
      this.onCorrectAnswer();
		} else {
      this.getAttempts();    
    }
	};

	onCorrectAnswer = () => {
    if(this.state.level === 4) {
      this.props.history.push('/dashboard')
    } else {
      this.getQuestion();
      this.getAttempts();
    }
	}


	openSnackbar = (message) => {
		this.setState({ barOpen: true, snackbarMessage: message });
	};

	onEditAnswer = e => {
		this.setState({ answer: e.target.value });
	};

	tries = () => {
		this.setState({
			tryvisible: true,
			hintvisible: false,
			statsvisible: false
		});
	};

	stats = () => {
		this.setState({
			tryvisible: false,
			hintvisible: false,
			statsvisible: true
		});
	};

	hints = () => {
		this.setState({
			tryvisible: false,
			hintvisible: true,
			statsvisible: false
		});
  };
  
	onKey = (e) => {
		if(e.key === 'Enter')
			this.checkAnswer();
	}

	render() {
		return (
			<div className={this.props.className}>
				{/* <Logout /> */}
				<h1 id="name">HAWKEYE</h1>
				<h2 id="region" />
				<div id="questionbox">
					<div id="level">{`Level ${this.state.level}`}</div>
					<div id="question">{this.state.question}</div>
          <span id="status">{this.state.hawkMessage}</span>
					<div id="answerbox">
						<input
							type="text"
							id="answer"
							placeholder="Enter answer here..."
							value={this.state.answer}
              onChange={this.onEditAnswer}
              onKeyPress={this.onKey}
						/>
					</div>
				</div>
				<Button onClick={this.checkAnswer} id="submit">
					SUBMIT
				</Button>
				<div id="hint_try">
					<div className="tab">
						<button
            style={{ width: window.innerWidth > 999 ? '100%' : '50%' }}
							className="tablinks"
							onClick={this.tries}
							id={this.state.tryvisible ? 'active' : 'inactive'}
						>
							Attempts
						</button>
						<button
            style={{ width: window.innerWidth > 999 ? '100%' : '50%' }}
							className="tablinks hintss"
							onClick={this.hints}
							id={this.state.hintvisible ? 'active' : 'inactive'}
						>
							Hints
						</button>
						{/* <button
							className="tablinks"
							onClick={this.stats}
							id={this.state.statsvisible ? 'active' : 'inactive'}
						>
							Stats
						</button> */}
					</div>
					<div
						id="hints"
						className={this.state.hintvisible ? 'available' : 'notavail'}
					>
						{this.state.hints.map(function(item, i) {
							return <p key={i}>{item}</p>;
						})}
					</div>

					<div
						id="tries"
						className={this.state.tryvisible ? 'available' : 'notavail'}
					>
						{this.state.attempts.map((item, i) => <p key={i}>{item}</p>)}
					</div>
					<div
						id="stats"
						className={this.state.statsvisible ? 'available' : 'notavail'}
					>
						{stats.map(function(item, i) {
							return <p key={i}>{item}</p>;
						})}
					</div>
				</div>

				<div id="hint_try1" className="hintnew">
					<div className="tab">
						<button className="tablinks" id="active">
							Hints
						</button>
					</div>
					<div id="hints" className="available">
						{this.state.hints.map(function(item, i) {
							return <p key={i}>{item}</p>;
						})}
					</div>
				</div>
				
				<div id="control">
					<div id="signals">
						<Link to="/rules"><i className="fas fa-question" /></Link>
						<Link to="/dashboard"><img src={logo} id="hawklogo" alt="" /></Link>
						<i className="fas fa-chess-rook" onClick={() => alert('Sidequest is currently locked. It will be accessible 12 hours into the game...')}/>
					</div>
					<img src={sideq} alt="" id="sideq" />
					<img src={map} alt="" id="map" />
				</div>
				<Snackbar open={this.state.barOpen} message={this.state.snackbarMessage}/>
			</div>
		);
	}
}

const drag = keyframes`
0%{
	opacity:0;
}
50%{
	opacity:0.8;
}
100%{
	opacity:1;
}
`;

export default styled(QuestionPage)`
      ${Logout}{
        z-index:150;
        position:absolute;
        left: 50%;
        top: 8%;
        transform: translate(-50%,50%);
      }
		#map {
			position: fixed;
			top: 0;
			z-index: -10;
			min-width: 100%;
			max-height: 200vh;
			opacity: 0.25;
		}
			#submit {
        left: 50%;
        top: calc(50vh + 170px);
        cursor: pointer;
        transform: translateX(-50%);
        z-index:30;
        position:absolute;
				color: #1c1c1c;
				font-weight: 500;
				background: #ffd627;
				width: 20%;
				max-width: 150px;
				height: 35px;
				padding: 10px;
				padding-top: 7px;
				border: none;
				border-radius: 20px;
				margin-top: 10px;
				margin-bottom: 10px;
			}
	.hintnew{
  display:none;
}
@media ${device.mobileS} {  
  #status{
    opacity:0.5;
    left: 50%;
        bottom: 25%;
        transform: translate(-50%,50%);
    position:absolute;
    z-index:150;
    color:white;
  }
  #submit{
    z-index:200;
    bottom:calc(50% - 230px);
    cursor: pointer;
    width:30%;
  }
  max-width: 420px; 
#name{
    text-align:center;
    font-size:2.5em;
			color:#FFD627;
			margin:10px 0 3px 0;
		}
  #questionbox{
        overflow:hidden;
        border-radius:10px;
		background: #242121;
		filter: drop-shadow(0px 15px 15px #000);
		width:80%;
		position: absolute;
		height:400px;
		text-align:center;
        left: 50%;
        top: 50%;
        transform: translate(-50%,-65%);
        z-index:100;

        #question{
            padding-top: 20px;
            font-size:1. 1.3rem;
            position:absolute;
            left: 50%;
        top: 10%;
        transform: translate(-50%,0%);
            width:80%;
             letter-spacing:1px;
				display:block;
				text-indent: 10px;
                font-weight:500;
                color:#fff;
        }

        #answerbox{
            position:absolute;
            bottom:0;
            height:60px;
            background:#FFD627;
            width: 100%;
						display: flex;
						flex-flow: row nowrap;
        }
        #answer{
            position:absolute;
            left: 50%;
        top: 10px;
        transform: translateX(-50%);
        width:80%;
        height:2.7em;
        font-size: 0.8em;
        letter-spacing:1px;
				display:block;
				text-indent: 10px;
				font-weight:500;
				border: 0px;
				border-radius: 4px;
				box-shadow: none;
				outline: none;
				-webkit-appearance:none ;
				-moz-appearance:none ;
                appearance:none ;
                background:rgba(255,255,255,0.3);
                text-transform: capitalize;
        }
  }
  #hint_try{
      overflow:hidden;
        border-radius:10px;
		background: #242121;
		filter: drop-shadow(0px 15px 15px #000);
		width:80%;
		position: absolute;
		height:320px;
		text-align:center;
        left: 50%;
        top: 50%;
        transform: translate(-50%,75%);
        z-index:100;
        margin-bottom:40px;
/* .tab {
  overflow: hidden;
  border: 1px solid #ccc;
  background-color: #f1f1f1; */
}

/* Style the buttons inside the tab */
.tab button {
    box-shadow: none;
				outline: none;
				-webkit-appearance:none ;
				-moz-appearance:none ;
                appearance:none ;
  border: none;
  outline: none;
  cursor: pointer;
  padding: 14px 16px;
  transition: 0.3s;
  font-size: 17px;
}
/* .tab button:hover {
  background-color: #ddd;
} */
.tab button{
  transition:0.4s;
    width:33.33%;
    font-size: 0.8em;
    letter-spacing:1px;
   font-weight:500;

}
.tabcontent {
  padding: 6px 12px;
  border: 1px solid #ccc;
  border-top: none;
}
.available{
    display :visible;
}
.notavail{
    display :none;
}

  }
#active{
  background-color: #FFD627;
  color:#000;

}

#inactive{
  background-color: transparent;
  color:#a2a3a2;
}

#tries{
  text-align:left;
  padding:10px;
  font-size:1.1em;
  color:#fff;
  letter-spacing:1px;
}
#hints{
  text-align:left;
  padding:10px;
  font-size:1.1em;
  color:#fff;
  letter-spacing:1px;
}
#stats{
  text-align:left;
  padding:10px;
  font-size:1.3em;
  color:#fff;
  letter-spacing:1px;
}

#bgq{
  z-index:-50;
  width:200%;
}
#sideq{
  z-index:101;
  position:absolute;
  bottom:0px;
  /* left: 50%;
        top: 50%;
        transform: translate(-50%,70%); */
  width:100%;
  filter: drop-shadow(0px -15px 10px #000);
}
#level{
  padding-bottom: 10px;
  border-bottom: solid thin rgba(100, 0, 0, 0.2);
  font-size:1.2em;
  font-weight:700;
  color:#fff;
  margin-top:10px;
}
#hawklogo{
  /* filter: drop-shadow(0px -15px 10px #000); */
  /* animation: ${drag} 2s 1 0s ease-in forwards; */
  width:15%;
  background:#181818;
  left: 50%;
  border:6px solid #c49905;
  padding:2px;
  border-radius:100px;
  position:absolute;
  bottom:20px;
  transform: translate(-50%,0%);
  z-index:102;
  transition:1s;
}
#hawklogo:hover{

  /* filter: drop-shadow(0px 2px 2px #000); */
}
#control{
  /* background:red; */
  position:relative;
  top:280px;
  height:100vh;
  
}
.fa-question{
  font-size:10vw;
  z-index:104;
  color:#242121;
 position:absolute;
  bottom:0.3vh;
  right:19%;
  /* transform: translate(-50%,0%); */
  padding:15px;
}
.fa-chess-rook{
  font-size:10vw;
  z-index:110;
  color:#242121;
position:absolute;
left:19%;
  bottom:0.3vh;
  /* transform: translate(-50%,0%); */
        padding:15px;
}

}
/******************  LARGE MOBILE  ************************/
@media ${device.mobileL} {
  max-width: 580px; 
#name{
    text-align:center;
    font-size:2.5em;
			color:#FFD627;
			margin:10px 0 3px 0;
		}
  #questionbox{
        overflow:hidden;
        border-radius:10px;
		background: #242121;
		filter: drop-shadow(0px 15px 15px #000);
		width:80%;
		position: absolute;
		height:400px;
		text-align:center;
        left: 50%;
        top: 50%;
        transform: translate(-50%,-65%);
        z-index:100;

        #question{
            padding-top: 20px;
            font-size: 1.3rem;
            position:absolute;
            left: 50%;
        top: 10%;
        transform: translate(-50%,0%);
            width:80%;
             letter-spacing:1px;
				display:block;
				text-indent: 10px;
                font-weight:500;
                color:#fff;
        }

        #answerbox{
            position:absolute;
            bottom:0;
            height:60px;
            background:#FFD627;
            width:100%;

        }
        #answer{
            position:absolute;
            left: 50%;
        top: 50%;
        transform: translate(-50%,-60%);
        width:80%;
        height:2.7em;
        font-size: 0.8em;
        letter-spacing:1px;
				display:block;
				text-indent: 10px;
				font-weight:500;
				border: 0px;
				border-radius: 4px;
				box-shadow: none;
				outline: none;
				-webkit-appearance:none ;
				-moz-appearance:none ;
                appearance:none ;
                background:rgba(255,255,255,0.3);
                text-transform: capitalize;
        }
  }
  #hint_try{
      overflow:hidden;
        border-radius:10px;
		background: #242121;
		filter: drop-shadow(0px 15px 15px #000);
		width:80%;
		position: absolute;
		height:320px;
		text-align:center;
        left: 50%;
        top: 50%;
        transform: translate(-50%, 80%);
        z-index:100;
        margin-bottom:40px;s

        .tab {
  overflow: hidden;
  border: 1px solid #ccc;
  background-color: #f1f1f1;
}

/* Style the buttons inside the tab */
.tab button {
    box-shadow: none;
				outline: none;
				-webkit-appearance:none ;
				-moz-appearance:none ;
                appearance:none ;
  border: none;
  outline: none;
  cursor: pointer;
  padding: 14px 16px;
  transition: 0.3s;
  font-size: 17px;
}
/* .tab button:hover {
  background-color: #ddd;
} */
.tab button{
  transition:0.4s;
    width:33.33%;
    font-size: 0.8em;
    letter-spacing:1px;
   font-weight:500;

}
.tabcontent {
  padding: 6px 12px;
  border: 1px solid #ccc;
  border-top: none;
}
.available{
    display :visible;
}
.notavail{
    display :none;
}

  }
#active{
  background-color: #FFD627;
  color:#000;

}

#inactive{
  background-color: transparent;
  color:#a2a3a2;
}

#tries{
  text-align:left;
  padding:10px;
  font-size:1.1em;
  color:#fff;
  letter-spacing:1px;
}
#hints{
  text-align:left;
  padding:10px;
  font-size:1.1em;
  color:#fff;
  letter-spacing:1px;
}
#stats{
  text-align:left;
  padding:10px;
  font-size:1.3em;
  color:#fff;
  letter-spacing:1px;
}

#bgq{
  z-index:-50;
  width:200%;
}
#sideq{
  z-index:101;
  position:absolute;
  bottom:0px;
  transform: translate(-50%,0);
  left: 50%;
  /* 
        top: 50%;
         */
  width:90%;
  filter: drop-shadow(0px -15px 10px #000);
}
#level{
  padding-bottom: 10px;
  border-bottom: solid thin rgba(100, 0, 0, 0.2);
  font-size:1.2em;
  font-weight:700;
  color:#fff;
  margin-top:10px;
}
#hawklogo{
  /* filter: drop-shadow(0px -15px 10px #000); */
  /* animation: ${drag} 2s 1 0s ease-in forwards; */
  width:14%;
  background:#181818;
  left: 50%;
  border:6px solid #c49905;
  padding:2px;
  border-radius:100px;
  position:absolute;
  bottom:20px;
  transform: translate(-50%,0%);
  z-index:102;
  transition:1s;
}
#hawklogo:hover{

  /* filter: drop-shadow(0px 2px 2px #000); */
}
#control{
  /* background:red; */
  position:relative;
  top:280px;
  height:100vh;
  
}
.fa-question{
  font-size:8vw;
  z-index:104;
  color:#242121;
 position:absolute;
  bottom:0.3vh;
  right:22%;
  /* transform: translate(-50%,0%); */
  padding:15px;
}
.fa-chess-rook{
  font-size:8vw;
  z-index:110;
  color:#242121;
position:absolute;
left:22%;
  bottom:0.3vh;
  /* transform: translate(-50%,0%); */
        padding:15px;
	}
}

/******************  TABLET  ************************/
@media ${device.tablet} {  
  max-width: 730px; 
#name{
    text-align:center;
    font-size:2.5em;
			color:#FFD627;
			margin:10px 0 3px 0;
		}
  #questionbox{
        overflow:hidden;
        border-radius:10px;
		background: #242121;
		filter: drop-shadow(0px 15px 15px #000);
		width:70%;
		position: absolute;
		height:400px;
		text-align:center;
        left: 50%;
        top: 50%;
        transform: translate(-50%,-65%);
        z-index:100;

        #question{
            padding-top: 20px;
            font-size: 1.3rem;
            position:absolute;
            left: 50%;
        top: 10%;
        transform: translate(-50%,0%);
            width:80%;
             letter-spacing:1px;
				display:block;
				text-indent: 10px;
                font-weight:500;
                color:#fff;
        }

        #answerbox{
            position:absolute;
            bottom:0;
            height:60px;
            background:#FFD627;
            width:100%;

        }
        #answer{
            position:absolute;
            left: 50%;
        top: 50%;
        transform: translate(-50%,-60%);
        width:80%;
        height:2.7em;
        font-size: 0.8em;
        letter-spacing:1px;
				display:block;
				text-indent: 10px;
				font-weight:500;
				border: 0px;
				border-radius: 4px;
				box-shadow: none;
				outline: none;
				-webkit-appearance:none ;
				-moz-appearance:none ;
                appearance:none ;
                background:rgba(255,255,255,0.3);
                text-transform: capitalize;
        }
  }
  #hint_try{
      overflow:hidden;
        border-radius:10px;
		background: #242121;
		filter: drop-shadow(0px 15px 15px #000);
		width:70%;
		position: absolute;
		height:320px;
		text-align:center;
        left: 50%;
        top: 50%;
        transform: translate(-50%, 80%);
        z-index:100;
        margin-bottom:40px;s

        .tab {
  overflow: hidden;
  border: 1px solid #ccc;
  background-color: #f1f1f1;
}

/* Style the buttons inside the tab */
.tab button {
    box-shadow: none;
				outline: none;
				-webkit-appearance:none ;
				-moz-appearance:none ;
                appearance:none ;
  border: none;
  outline: none;
  cursor: pointer;
  padding: 14px 16px;
  transition: 0.3s;
  font-size: 17px;
}
/* .tab button:hover {
  background-color: #ddd;
} */
.tab button{
  transition:0.4s;
    width:33.33%;
    font-size: 0.9em;
    letter-spacing:1px;
   font-weight:500;

}
.tabcontent {
  padding: 6px 12px;
  border: 1px solid #ccc;
  border-top: none;
}
.available{
    display :visible;
}
.notavail{
    display :none;
}

  }
#active{
  background-color: #FFD627;
  color:#000;

}

#inactive{
  background-color: transparent;
  color:#a2a3a2;
}

#tries{
  text-align:left;
  padding:10px;
  font-size:1.1em;
  color:#fff;
  letter-spacing:1px;
}
#hints{
  text-align:left;
  padding:10px;
  font-size:1.1em;
  color:#fff;
  letter-spacing:1px;
}
#stats{
  text-align:left;
  padding:10px;
  font-size:1.3em;
  color:#fff;
  letter-spacing:1px;
}

#bgq{
  z-index:-50;
  width:200%;
}
#sideq{
  z-index:101;
  position:absolute;
  bottom:0px;
  transform: translate(-50%,0);
  left: 50%;
  /* 
        top: 50%;
         */
  width:75%;
  filter: drop-shadow(0px -15px 10px #000);
}
#level{
  padding-bottom: 10px;
  border-bottom: solid thin rgba(100, 0, 0, 0.2);
  font-size:1.2em;
  font-weight:700;
  color:#fff;
  margin-top:10px;
}
#hawklogo{
  /* filter: drop-shadow(0px -15px 10px #000); */
  /* animation: ${drag} 2s 1 0s ease-in forwards; */
  width:11%;
  background:#181818;
  left: 50%;
  border:6px solid #c49905;
  padding:2px;
  border-radius:100px;
  position:absolute;
  bottom:20px;
  transform: translate(-50%,0%);
  z-index:102;
  transition:1s;
}
#hawklogo:hover{

  /* filter: drop-shadow(0px 2px 2px #000); */
}
#control{
  /* background:red; */
  position:relative;
  top:280px;
  height:100vh;
  
}
.fa-question{
  font-size:6vw;
  z-index:104;
  color:#242121;
 position:absolute;
  bottom:0.3vh;
  right:26%;
  /* transform: translate(-50%,0%); */
  padding:15px;
}
.fa-chess-rook{
  font-size:6vw;
  z-index:110;
  color:#242121;
position:absolute;
left:26%;
  bottom:0.3vh;
  /* transform: translate(-50%,0%); */
        padding:15px;
}
}












/******************  TABLETlarge  ************************/
@media ${device.laptop} {  
  max-width: 862px; 
#name{
    text-align:center;
    font-size:2.5em;
			color:#FFD627;
			margin:10px 0 3px 0;
		}
  #questionbox{
        overflow:hidden;
        border-radius:10px;
		background: #242121;
		filter: drop-shadow(0px 15px 15px #000);
		width:70%;
		position: absolute;
		height:400px;
		text-align:center;
        left: 50%;
        top: 50%;
        transform: translate(-50%,-65%);
        z-index:100;

        #question{
            padding-top: 20px;
            font-size: 1.3rem;
            position:absolute;
            left: 50%;
        top: 10%;
        transform: translate(-50%,0%);
            width:80%;
             letter-spacing:1px;
				display:block;
				text-indent: 10px;
                font-weight:500;
                color:#fff;
        }

        #answerbox{
            position:absolute;
            bottom:0;
            height:60px;
            background:#FFD627;
            width:100%;

        }
        #answer{
            position:absolute;
            left: 50%;
        top: 50%;
        transform: translate(-50%,-60%);
        width:80%;
        height:2.7em;
        font-size: 0.8em;
        letter-spacing:1px;
				display:block;
				text-indent: 10px;
				font-weight:500;
				border: 0px;
				border-radius: 4px;
				box-shadow: none;
				outline: none;
				-webkit-appearance:none ;
				-moz-appearance:none ;
                appearance:none ;
                background:rgba(255,255,255,0.3);
                text-transform: capitalize;
        }
  }
  #hint_try{
      overflow:hidden;
        border-radius:10px;
		background: #242121;
		filter: drop-shadow(0px 15px 15px #000);
		width:70%;
		position: absolute;
		height:320px;
		text-align:center;
        left: 50%;
        top: 50%;
        transform: translate(-50%, 80%);
        z-index:100;
        margin-bottom:40px;s

        .tab {
  overflow: hidden;
  border: 1px solid #ccc;
  background-color: #f1f1f1;
}

/* Style the buttons inside the tab */
.tab button {
    box-shadow: none;
				outline: none;
				-webkit-appearance:none ;
				-moz-appearance:none ;
                appearance:none ;
  border: none;
  outline: none;
  cursor: pointer;
  padding: 14px 16px;
  transition: 0.3s;
  font-size: 17px;
}
/* .tab button:hover {
  background-color: #ddd;
} */
.tab button{
  transition:0.4s;
    width:33.33%;
    font-size: 0.9em;
    letter-spacing:1px;
   font-weight:500;

}
.tabcontent {
  padding: 6px 12px;
  border: 1px solid #ccc;
  border-top: none;
}
.available{
    display :visible;
}
.notavail{
    display :none;
}

  }
#active{
  background-color: #FFD627;
  color:#000;

}

#inactive{
  background-color: transparent;
  color:#a2a3a2;
}

#tries{
  text-align:left;
  padding:10px;
  font-size:1.1em;
  color:#fff;
  letter-spacing:1px;
}
#hints{
  text-align:left;
  padding:10px;
  font-size:1.1em;
  color:#fff;
  letter-spacing:1px;
}
#stats{
  text-align:left;
  padding:10px;
  font-size:1.3em;
  color:#fff;
  letter-spacing:1px;
}

#bgq{
  z-index:-50;
  width:200%;
}
#sideq{
  z-index:101;
  position:absolute;
  bottom:0px;
  transform: translate(-50%,0);
  left: 50%;
  /* 
        top: 50%;
         */
  width:65%;
  filter: drop-shadow(0px -15px 10px #000);
}
#level{
  padding-bottom: 10px;
  border-bottom: solid thin rgba(100, 0, 0, 0.2);
  font-size:1.2em;
  font-weight:700;
  color:#fff;
  margin-top:10px;
}
#hawklogo{
  /* filter: drop-shadow(0px -15px 10px #000); */
  /* animation: ${drag} 2s 1 0s ease-in forwards; */
  width:9%;
  background:#181818;
  left: 50%;
  border:6px solid #c49905;
  padding:2px;
  border-radius:100px;
  position:absolute;
  bottom:20px;
  transform: translate(-50%,0%);
  z-index:102;
  transition:1s;
}
#hawklogo:hover{

  /* filter: drop-shadow(0px 2px 2px #000); */
}
#control{
  /* background:red; */
  position:relative;
  top:280px;
  height:100vh;
  
}
.fa-question{
  font-size:6vw;
  z-index:104;
  color:#242121;
 position:absolute;
  bottom:0.3vh;
  right:29%;
  /* transform: translate(-50%,0%); */
  padding:15px;
}
.fa-chess-rook{
  font-size:6vw;
  z-index:110;
  color:#242121;
position:absolute;
left:29%;
  bottom:0.3vh;
  /* transform: translate(-50%,0%); */
        padding:15px;
}
}








/******************  laptop  ************************/
@media ${device.laptopL} {  
  max-width: 1000px; 
#name{
    text-align:center;
    font-size:2.5em;
			color:#FFD627;
			margin:10px 0 3px 0;
		}
  #questionbox{
        overflow:hidden;
        border-radius:10px;
		background: #242121;
		filter: drop-shadow(0px 15px 15px #000);
		width:50%;
		position: absolute;
		height:400px;
		text-align:center;
        left: 50%;
        top: 50%;
        transform: translate(-50%,-65%);
        z-index:100;

        #question{
            padding-top: 20px;
            font-size: 1.3rem;
            position:absolute;
            left: 50%;
        top: 10%;
        transform: translate(-50%,0%);
            width:80%;
             letter-spacing:1px;
				display:block;
				text-indent: 10px;
                font-weight:500;
                color:#fff;
        }

        #answerbox{
            position:absolute;
            bottom:0;
            height:60px;
            background:#FFD627;
            width:100%;

        }
        #answer{
            position:absolute;
            left: 50%;
        top: 50%;
        transform: translate(-50%,-60%);
        width:80%;
        height:2.7em;
        font-size: 0.8em;
        letter-spacing:1px;
				display:block;
				text-indent: 10px;
				font-weight:500;
				border: 0px;
				border-radius: 4px;
				box-shadow: none;
				outline: none;
				-webkit-appearance:none ;
				-moz-appearance:none ;
                appearance:none ;
                background:rgba(255,255,255,0.3);
                text-transform: capitalize;
        }
  }
  #hint_try{
      overflow:hidden;
        border-radius:10px;
		background: #242121;
		filter: drop-shadow(0px 15px 15px #000);
		width:50%;
		position: absolute;
		height:320px;
		text-align:center;
        left: 50%;
        top: 50%;
        transform: translate(-50%, 80%);
        z-index:100;
        margin-bottom:40px;s

        .tab {
  overflow: hidden;
  border: 1px solid #ccc;
  background-color: #f1f1f1;
}

/* Style the buttons inside the tab */
.tab button {
    box-shadow: none;
				outline: none;
				-webkit-appearance:none ;
				-moz-appearance:none ;
                appearance:none ;
  border: none;
  outline: none;
  cursor: pointer;
  padding: 14px 16px;
  transition: 0.3s;
  font-size: 17px;
}
/* .tab button:hover {
  background-color: #ddd;
} */
.tab button{
  transition:0.4s;
    width:33.33%;
    font-size: 0.9em;
    letter-spacing:1px;
   font-weight:500;

}
.tabcontent {
  padding: 6px 12px;
  border: 1px solid #ccc;
  border-top: none;
}
.available{
    display :visible;
}
.notavail{
    display :none;
}

  }
#active{
  background-color: #FFD627;
  color:#000;

}

#inactive{
  background-color: transparent;
  color:#a2a3a2;
}

#tries{
  text-align:left;
  padding:10px;
  font-size:1.1em;
  color:#fff;
  letter-spacing:1px;
}
#hints{
  text-align:left;
  padding:10px;
  font-size:1.1em;
  color:#fff;
  letter-spacing:1px;
}
#stats{
  text-align:left;
  padding:10px;
  font-size:1.3em;
  color:#fff;
  letter-spacing:1px;
}

#bgq{
  z-index:-50;
  width:200%;
}
#sideq{
  z-index:101;
  position:absolute;
  bottom:0px;
  transform: translate(-50%,0);
  left: 50%;
  /* 
        top: 50%;
         */
  width:45%;
  filter: drop-shadow(0px -15px 10px #000);
}
#level{
  padding-bottom: 10px;
  border-bottom: solid thin rgba(100, 0, 0, 0.2);
  font-size:1.2em;
  font-weight:700;
  color:#fff;
  margin-top:10px;
}
#hawklogo{
  /* filter: drop-shadow(0px -15px 10px #000); */
  /* animation: ${drag} 2s 1 0s ease-in forwards; */
  width:7%;
  background:#181818;
  left: 50%;
  border:6px solid #c49905;
  padding:2px;
  border-radius:100px;
  position:absolute;
  bottom:20px;
  transform: translate(-50%,0%);
  z-index:102;
  transition:1s;
}
#hawklogo:hover{

  /* filter: drop-shadow(0px 2px 2px #000); */
}
#control{
  /* background:red; */
  position:relative;
  top:280px;
  height:100vh;
  
}
.fa-question{
  font-size:4vw;
  z-index:104;
  color:#242121;
 position:absolute;
  bottom:0.3vh;
  right:35%;
  /* transform: translate(-50%,0%); */
  padding:15px;
}
.fa-chess-rook{
  font-size:4vw;
  z-index:110;
  color:#242121;
position:absolute;
left:35%;
  bottom:0.3vh;
  /* transform: translate(-50%,0%); */
        padding:15px;
}

}

















/******************  larggetsssss ************************/
@media ${device.desktop} { 
		#map {
			position: fixed;
			top: 0;
			min-width: 100%;
			max-height: 400vh;
			opacity: 0.25;
		}
  .hintnew{
  display:block;
}
  max-width: 2000px; 
#name{
    text-align:center;
    font-size:2.5em;
			color:#FFD627;
			margin:10px 0 3px 0;
		}
  #questionbox{
        overflow:hidden;
        border-radius:10px;
		background: #242121;
		filter: drop-shadow(0px 15px 15px #000);
		width:35%;
		position: absolute;
		height:380px;
		text-align:center;
        left: 50%;
        top: 50%;
        transform: translate(-50%,-55%);
        z-index:100;

        #question{
            padding-top: 20px;
            font-size:2. 1.3rem;
            position:absolute;
            left: 50%;
        top: 10%;
        transform: translate(-50%,0%);
            width:80%;
             letter-spacing:1px;
				display:block;
				text-indent: 10px;
                font-weight:500;
                color:#fff;
        }

        #answerbox{
            position:absolute;
            bottom:0;
            height:60px;
            background:#FFD627;
            width:100%;

        }
        #answer{
            position:absolute;
            left: 50%;
        top: 50%;
        transform: translate(-50%,-60%);
        width:80%;
        height:2.7em;
        font-size: 0.8em;
        letter-spacing:1px;
				display:block;
				text-indent: 10px;
				font-weight:500;
				border: 0px;
				border-radius: 4px;
				box-shadow: none;
				outline: none;
				-webkit-appearance:none ;
				-moz-appearance:none ;
                appearance:none ;
                background:rgba(255,255,255,0.3);
                text-transform: capitalize;
        }
  }
  #hint_try{
      overflow:hidden;
        border-radius:10px;
		background: #242121;
		filter: drop-shadow(0px 15px 15px #000);
		width:25%;
		position: absolute;
		height:320px;
		text-align:center;
        left: 2%;
        top: 50%;
        transform: translate(0%,-55%);
        z-index:100;
        margin-bottom:40px;s

        .tab {
  overflow: hidden;
  border: 1px solid #ccc;
  background-color: #f1f1f1;
}

/* Style the buttons inside the tab */
.tab button {
    box-shadow: none;
				outline: none;
				-webkit-appearance:none ;
				-moz-appearance:none ;
                appearance:none ;
  border: none;
  outline: none;
  cursor: pointer;
  padding: 14px 16px;
  transition: 0.3s;
  font-size: 14px;
}
/* .tab button:hover {
  background-color: #ddd;
} */
.tab button{
  transition:0.4s;
    width:50%;
    font-size: 0.9em;
    letter-spacing:1px;
   font-weight:500;

}
.tabcontent {
  padding: 6px 12px;
  border: 1px solid #ccc;
  border-top: none;
}
.available{
    display :visible;
}
.notavail{
    display :none;
}

  }





#hint_try1{
      overflow:hidden;
        border-radius:10px;
		background: #242121;
		filter: drop-shadow(0px 15px 15px #000);
		width:25%;
		position: absolute;
		height:320px;
		text-align:center;
        right: 2%;
        top: 50%;
        transform: translate(0%,-55%);
        z-index:100;
        margin-bottom:40px;s

        .tab {
  overflow: hidden;
  border: 1px solid #ccc;
  background-color: #f1f1f1;
}

/* Style the buttons inside the tab */
.tab button {
    box-shadow: none;
				outline: none;
				-webkit-appearance:none ;
				-moz-appearance:none ;
                appearance:none ;
  border: none;
  outline: none;
  cursor: pointer;
  padding: 14px 16px;
  transition: 0.3s;
  font-size: 14px;
}
/* .tab button:hover {
  background-color: #ddd;
} */
.tab button{
  transition:0.4s;
    width:100%;
    font-size: 0.9em;
    letter-spacing:1px;
   font-weight:500;

}
.tabcontent {
  padding: 6px 12px;
  border: 1px solid #ccc;
  border-top: none;
}
.available{
    display :visible;
}
.notavail{
    display :none;
}

  }





#active{
  background-color: #FFD627;
  color:#000;

}

#inactive{
  background-color: transparent;
  color:#a2a3a2;
}

#tries{
  text-align:left;
  padding:10px;
  font-size:1em;
  color:#fff;
  letter-spacing:1px;
}
#hints{
  text-align:left;
  padding:10px;
  font-size:1em;
  color:#fff;
  letter-spacing:1px;
}

#stats{
  text-align:left;
  padding:10px;
  font-size:1em;
  color:#fff;
  letter-spacing:1px;
}

#bgq{
  z-index:-50;
  width:200%;
}
#sideq{
  z-index:20;
  position:fixed;
  bottom:0px;
  transform: translate(-50%,0);
  left: 50%;
  /* 
        top: 50%;
         */
  width:30%;
  filter: drop-shadow(0px -5px 10px #000);
}
#level{
  padding-bottom: 10px;
  border-bottom: solid thin rgba(100, 100, 100, 0.2);
  font-size:1.2em;
  font-weight:700;
  color:#fff;
  margin-top:10px;
}
#hawklogo{
  /* filter: drop-shadow(0px -15px 10px #000); */
  /* animation: ${drag} 2s 1 0s ease-in forwards; */
  width: 70px;
  background:#181818;
  left: 50%;
  border:5px solid #c49905;
  padding:2px;
  border-radius:100px;
  position:fixed;
  bottom:20px;
  transform: translate(-50%,0%);
  z-index:102;
  transition:0.3s;
}
#hawklogo:hover{
  width: 80px;
  /* filter: drop-shadow(0px 2px 2px #000); */
}
#sideq {
  z-index: 20;
  position: fixed;
  bottom: 0px;
  transform: translate(-50%, 0);
  left: 50%;
  /* 
    top: 50%;
     */
  width: 400px;
  filter: drop-shadow(0px -5px 10px #000);
}
#control {
  position: fixed;
  bottom: 0;
  /* background:red; */
  position: relative;
  top: 0px;
  height: 18vh;
}
.fa-question {
  font-size: 30px;
  z-index: 100;
  color: #242121;
  position: fixed;
  bottom: 0px;
  right: calc(50% - 125px);
  transition: 0.3s;
  /* transform: translate(-50%,0%); */
  padding: 20px;
}
.fa-chess-rook {
  font-size: 30px;
  z-index: 100;
  color: #242121;
  position: fixed;
  left: calc(50% - 125px);
  bottom: 0px;
  transition: 0.3s;
  /* transform: translate(-50%,0%); */
  padding: 20px;
}
.fa-chess-rook:hover,
.fa-question:hover {
  cursor: pointer;
  font-size: 35px;
}

#hawklogo:hover{
  cursor:pointer;
  font-size:3.2vw;
}

.hintss{
  display:none;
}
}
`;
