import React, { Component } from 'react';
import styled from 'styled-components';
import { observer, inject } from 'mobx-react';
import { UserStore } from '../stores/User';
import media from '../components/theme/media';
import { centerHV } from '../mixins';
import * as colors from '../components/';
import map from '../components/assets/mapbg.png';
import logo from '../components/assets/iecse_logo.png';
import hawk from '../components/assets/hawk_logo.png';

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
import { Snackbar } from '../components';
import { Link, Redirect } from 'react-router-dom';

import { Button, TextField } from '../components';

type Props = {
	className: string;
	UserStore: UserStore;
};

@inject('UserStore')
@observer
class LoginPage extends Component<Props> {
	state = {
		username: '',
		password: '',
		loggedIn: false,
		barOpen: false,
		snackbarMessage: '',
		sending: false,
		firstlog: 1
	};

	onChange = (name, value) => {
		this.setState({ [name]: value });
	};

	onKey = e => {
		if (e.key === 'Enter') this.login();
	};

	login = () => {
		if (this.state.username === '' || this.state.password === '') return;
		const { loggedIn, ...loginData } = this.state;
		this.setState({ sending: true }, () => {
			fetch('/api/login', {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify(loginData)
			})
				.then(res => res.json())
				.then(json => {
					this.setState({ sending: false });
					if (json.success) {
						this.props.UserStore.setCurrentUser(json.data);
						this.setState({ loggedIn: true, firstlog: json.data.firstLogin });
					} else {
						this.openSnackbar(json.msg);
					}
				});
		});
	};

	openSnackbar = message => {
		this.setState({ barOpen: true, snackbarMessage: message });
		setTimeout(
			() => this.setState({ barOpen: false, snackbarMessage: '' }),
			3000
		);
	};

	render() {
		return (
			<div className={this.props.className}>
				{/* <img src=hawk} alt="" id="hawk" /> */}
				<div id="box">
					<h1>HAWKEYE</h1>
					<span id="subt">Sign into your account</span>

					<div id="inputs">
						<TextField
							name="username"
							placeholder="Username"
							onChange={this.onChange}
							onKeyPress={this.onKey}
						/>
						<br />
						<TextField
							name="password"
							type="password"
							placeholder="Password"
							onChange={this.onChange}
							onKeyPress={this.onKey}
						/>
						<br />
					</div>
					<Button disabled={this.state.sending} onClick={this.login}>
						<i
							className="fa fa-spinner fa-spin"
							id={this.state.sending ? 'loading' : 'notloading'}
						/>
						{this.state.sending ? ' Sending...' : 'Sign In'}
					</Button>
					{this.state.loggedIn ? (
						this.state.firstlog ? (
							<Redirect to="/tutorial" />
						) : (
							<Redirect to="/dashboard" />
						)
					) : null}
					<br />
					<Link to="/register" id="register">
						Create an account
					</Link>
				</div>
				<img src={map} alt="" id="map" />
				<a href="https://www.iecsemanipal.com/">
					<img src={logo} alt="" id="logo" />
				</a>
				<Snackbar
					open={this.state.barOpen}
					message={this.state.snackbarMessage}
				/>
			</div>
		);
	}
}

export default styled(LoginPage)`
	#loading {
		display: inline-block;
	}
	#notloading {
		display: none;
	}
	width: 100%;
	@media ${device.mobileS} {
		max-width: 500px;
		#box {
			padding-top: 10px;
			h1 {
				color: #ffd627;
				margin: 10px 0 3px 0;
			}
			span {
				color: white;
				font-weight: 300;
				font-size: 0.7em;
				letter-spacing: 3px;
			}
			border-radius: 10px;

			background: #1c1c1c;
			filter: drop-shadow(0px 15px 15px #000);
			width: 80%;
			position: absolute;
			height: 280px;
			text-align: center;
			left: 50%;
			top: 50%;
			transform: translate(-50%, -50%);
			#inputs {
				${TextField} {
					input {
						letter-spacing: 1.25px;
						font-size: 100%;
						text-indent: 10px;
						margin-top: 10px;
						font-weight: 300;
						width: 80%;
						color: #fff;
						border: 0px;
						border-radius: 4px;
						box-shadow: none;
						outline: none;
						-webkit-appearance: none;
						-moz-appearance: none;
						appearance: none;
						height: 25px;
						border: none;
						text-indent: 10px;
						background: #333333;
					}
				}
			}

			${Button} {
				color: #1c1c1c;
				font-weight: 500;
				background: #ffd627;
				width: 50%;
				height: 35px;
				padding: 10px;
				padding-top: 7px;
				border: none;
				border-radius: 20px;
				margin-top: 10px;
				margin-bottom: 10px;
			}
			#register {
				letter-spacing: 1px;
				color: white;
				text-decoration: none;
				font-weight: 300;
				font-size: 0.8em;
			}
		}

		#map {
			position: fixed;
			top: 0;
			z-index: -10;
			min-width: 100%;
			max-height: 200vh;
			opacity: 0.25;
		}
		#logo {
			position: absolute;
			width: 10%;
			bottom: 30px;
			left: 50%;
			transform: translate(-50%, 0%);
		}
	}

	/************* mobmedium *************/
	@media ${device.mobileL} {
		max-width: 768px;
		#box {
			padding-top: 10px;
			h1 {
				color: #ffd627;
				margin: 10px 0 3px 0;
			}
			span {
				color: white;
				font-weight: 300;
				font-size: 0.7em;
				letter-spacing: 3px;
			}
			border-radius: 10px;

			background: #1c1c1c;
			filter: drop-shadow(0px 15px 15px #000);
			width: 60%;
			position: absolute;
			height: 280px;
			text-align: center;
			left: 50%;
			top: 50%;
			transform: translate(-50%, -50%);
			#inputs {
				${TextField} {
					input {
						letter-spacing: 1.25px;
						font-size: 100%;
						text-indent: 10px;
						margin-top: 10px;
						font-weight: 300;
						width: 80%;
						color: #fff;
						border: 0px;
						border-radius: 4px;
						box-shadow: none;
						outline: none;
						-webkit-appearance: none;
						-moz-appearance: none;
						appearance: none;
						height: 25px;
						border: none;
						text-indent: 10px;
						background: #333333;
					}
				}
			}

			${Button} {
				color: #1c1c1c;
				font-weight: 500;
				background: #ffd627;
				width: 50%;
				height: 35px;
				padding: 10px;
				padding-top: 7px;
				border: none;
				border-radius: 20px;
				margin-top: 10px;
				margin-bottom: 10px;
			}
			#register {
				letter-spacing: 1px;
				color: white;
				text-decoration: none;
				font-weight: 300;
				font-size: 0.8em;
			}
		}

		#map {
			position: fixed;
			top: 0;
			min-width: 100%;
			max-height: 200vh;
			opacity: 0.25;
			/* Preserve aspet ratio */

			/* position: absolute;
			z-index: -10;
			width: 100vw;
			top: -100vh;
			/* left: 50%; */

			/* transform: translate(-50%, 0%); */
		}
		#logo {
			position: absolute;
			width: 10%;
			bottom: 30px;
			left: 50%;
			transform: translate(-50%, 0%);
		}
	}

	/************* tablet *************/
	@media ${device.tablet} {
		max-width: 1000px;
		#box {
			padding-top: 10px;
			h1 {
				color: #ffd627;
				margin: 10px 0 3px 0;
			}
			span {
				color: white;
				font-weight: 300;
				font-size: 0.7em;
				letter-spacing: 3px;
			}
			border-radius: 10px;

			background: #1c1c1c;
			filter: drop-shadow(0px 15px 15px #000);
			width: 50%;
			position: absolute;
			height: 300px;
			text-align: center;
			left: 50%;
			top: 50%;
			transform: translate(-50%, -50%);
			#inputs {
				${TextField} {
					input {
						letter-spacing: 1.25px;
						font-size: 100%;
						text-indent: 10px;
						margin-top: 10px;
						font-weight: 300;
						width: 80%;
						color: #fff;
						border: 0px;
						border-radius: 4px;
						box-shadow: none;
						outline: none;
						-webkit-appearance: none;
						-moz-appearance: none;
						appearance: none;
						height: 25px;
						border: none;
						text-indent: 10px;
						background: #333333;
					}
				}
			}

			${Button} {
				color: #1c1c1c;
				font-weight: 500;
				background: #ffd627;
				width: 50%;
				height: 35px;
				padding: 10px;
				padding-top: 7px;
				border: none;
				border-radius: 20px;
				margin-top: 10px;
				margin-bottom: 10px;
			}
			#register {
				letter-spacing: 1px;
				color: white;
				text-decoration: none;
				font-weight: 300;
				font-size: 0.8em;
			}
		}

		#map {
			position: fixed;
			top: 0;
			min-width: 100%;
			max-height: 200vh;
			opacity: 0.25;
		}
		#logo {
			position: absolute;
			width: 5%;
			bottom: 30px;
			left: 50%;
			transform: translate(-50%, 0%);
		}
	}

	/************* desktop *************/
	@media ${device.desktop} {
		#box {
			padding-top: 10px;
			h1 {
				color: #ffd627;
				margin: 10px 0 3px 0;
			}
			span {
				color: white;
				font-weight: 300;
				font-size: 0.7em;
				letter-spacing: 3px;
			}
			border-radius: 10px;

			background: #1c1c1c;
			filter: drop-shadow(0px 15px 15px #000);
			width: 30%;
			position: absolute;
			height: 350px;
			text-align: center;
			left: 50%;
			top: 50%;
			transform: translate(-50%, -50%);
			#inputs {
				${TextField} {
					input {
						letter-spacing: 1.25px;
						font-size: 100%;
						text-indent: 10px;
						margin-top: 10px;
						font-weight: 300;
						width: 80%;
						color: #fff;
						border: 0px;
						border-radius: 4px;
						box-shadow: none;
						outline: none;
						-webkit-appearance: none;
						-moz-appearance: none;
						appearance: none;
						height: 25px;
						border: none;
						text-indent: 10px;
						background: #333333;
					}
				}
			}

			${Button} {
				color: #1c1c1c;
				font-weight: 500;
				background: #ffd627;
				width: 50%;
				height: 35px;
				padding: 10px;
				padding-top: 7px;
				border: none;
				border-radius: 20px;
				margin-top: 10px;
				margin-bottom: 10px;
				transition: 0.5s;
			}
			${Button}:hover {
				color: #fff;
				background: #ff0000;
				width: 60%;
			}
			#register {
				letter-spacing: 1px;
				color: white;
				text-decoration: none;
				font-weight: 300;
				font-size: 0.8em;
			}
		}

		#map {
			position: fixed;
			top: 0;
			min-width: 100%;
			max-height: 300vh;
			opacity: 0.25;
		}
		#logo {
			position: absolute;
			width: 5%;
			bottom: 30px;
			left: 50%;
			transform: translate(-50%, 0%);
		}
	}
`;
