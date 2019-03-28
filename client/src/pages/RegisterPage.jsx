import React, { Component } from 'react';
import PropTypes from 'prop-types';
import styled, { keyframes } from 'styled-components';
import validator from 'validator';
import Recaptcha from 'react-recaptcha';
import media from '../components/theme/media';
import logo from '../components/assets/iecse_logo.png';
import hawk from '../components/assets/hawk_logo.png';
import pcb from '../components/assets/pcbdesign.png';
import pcb1 from '../components/assets/pcbdesign1.png';
import { Snackbar } from '../components';
import { Redirect, Link } from 'react-router-dom';

import { Button, TextField } from '../components';

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

class RegisterPage extends Component {
	state = {
		formData: {
			name: '',
			username: '',
			password: '',
			confirm_password: '',
			email: '',
			tel: '',
			college: '',
			token: ''
		},
		isVerified: false,
		snackbarMessage: '',
		barOpen: false,
		redirect: false
	};

	onChange = (name, value, error) => {
		this.setState({
			formData: Object.assign(this.state.formData, { [name]: { value, error } })
		});
	};

	recaptchaLoaded = () => {
		if (this.captchaDemo) {
			this.captchaDemo.reset();
		}
	};

	verifyCallback = token => {
		if (token) {
			this.setState({ isVerified: true });
			this.setState({ token });
		}
	};

	onSubmit = () => {
		// checkerror();
		if (this.state.isVerified) {
			// Remove error keys from formData
			const postData = Object.entries(this.state.formData)
				.map(([k, v]) => {
					return { [k]: v.value };
				})
				.reduce((acc, val) => Object.assign(acc, val));

			fetch('/api/addUser', {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify({ user: postData, captcha: this.state.token})
			})
				.then(res => res.json())
				.then(json => { 
					if (json.success) {
						this.openSnackbar('Registered successfully')
						this.setState({
							redirect: true
						});
					} else {
						this.openSnackbar('Registration failed')
					} 
				})
				.catch(() => this.openSnackbar('Registration failed'));
		} else {
			this.openSnackbar('Please verify that you are a human')
		}
	};

	openSnackbar = (message) => {
		this.setState({ barOpen: true, snackbarMessage: message });
		setTimeout(() => this.setState({ barOpen: false, snackbarMessage: '' }), 1000);
	};

	render() {

		return (
			<div className={this.props.className}>
				<div id="box">
					<h1>HAWKEYE</h1>
					<span>Create a new account</span>
					<div id="inputs">
						<TextField
							name="name"
							placeholder="Name"
							onChange={this.onChange}
						/>
						<TextField
							name="username"
							placeholder="Username"
							onChange={this.onChange}
						/>
						<TextField
							name="password"
							type="password"
							placeholder="Password"
							onChange={this.onChange}
						/>
						<TextField
							name="confirm_password"
							type="password"
							placeholder="Confirm password"
							onChange={this.onChange}
						/>
						<TextField
							name="email"
							type="email"
							placeholder="Email"
							onChange={this.onChange}
							validation={v => (validator.isEmail(v) ? '' : 'Invalid Email')}
							validateOnChange
						/>
						<TextField
							name="tel"
							placeholder="Mobile Number"
							onChange={this.onChange}
							validation={v =>
								validator.isMobilePhone(v) ? '' : 'Invalid Number'
							}
							validateOnChange
						/>
						<TextField
							name="college"
							placeholder="College"
							onChange={this.onChange}
						/>
					</div>
					<Recaptcha
						id="recap"
						ref={el => {
							this.captchaDemo = el;
						}}
						size="normal"
						theme="dark"
						render="explicit"
						sitekey="6LftZZoUAAAAAJPGqjgVjZ0ZI9aPQ7RJWKvocH1g"
						onloadCallback={this.recaptchaLoaded}
						verifyCallback={this.verifyCallback}
					/>
					<Button onClick={this.onSubmit} id="regbtn">
						Register
					</Button>
					<span><br/><Link to="/" style={{ 'color': 'white', 'text-decoration': 'none', 'font-weight': '500', 'font-size': '1.2em'}}>Already have an account? Log in </Link></span>
				</div>

				<img src={hawk} alt="" id="hawk" />
				<img src={logo} alt="" id="logo" />
				<div id="bg">
					<img src={pcb} alt="" id="pcb1" />
					<img src={pcb} alt="" id="pcb2" />
					<img src={pcb1} alt="" id="pcbdesk" />
				</div>
			<Snackbar open={this.state.barOpen} message={this.state.snackbarMessage}/>
			{ this.state.redirect && <Redirect to="/"/>}
			</div>
	);
	}
}

RegisterPage.propTypes = {
	className: PropTypes.string
};

const flash = keyframes`
0%{
	opacity:0.1;
}
50%{
	opacity:0.3;
}
100%{
	opacity:0.1
}
`;

export default styled(RegisterPage)`
	/* #recap.rc-anchor-dark.rc-anchor-normal {
		border: none !important;
	}*/
	#g-recaptcha {
		margin: 0 auto;
		display: table;
	}

	#pcbdesk {
		display: none;
	}
	@media ${device.mobileS} {
		max-width: 768px;
		#hawk {
			position: absolute;
			z-index: -10;
			left: 10px;
			top: 10px;
			/* transform: translate(-50%,0%); */
			width: 10%;
		}
		#logo {
			position: absolute;
			z-index: -10;
			right: 10px;
			top: 16px;
			/* transform: translate(-50%,0%); */
			width: 8%;
		}
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
			height: 520px;
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
					.error {
						margin-left: auto;
						margin-right: auto;
						margin-top: 10px;
						text-align: center;
						width: 34%;
						padding: 5px;
						/* margin: 10px; */
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
		}

		#pcb1 {
			left: 50%;
			top: 0%;
			transform: translate(-50%, 0%);
			width: 100%;
			position: relative;
		}
		#pcb2 {
			left: 50%;
			bottom: 10px;
			transform: translate(-50%, 0%);
			width: 100%;
			position: relative;
		}
		#bg {
			animation: ${flash} 2s infinite 0s ease-in-out;
			opacity: 0.1;
			z-index: -100;
			position: absolute;
			overflow-y: hidden;
			left: 50%;
			top: 50%;
			transform: translate(-50%, -50%);
			width: 100vw;
			height: 100%;
			/* background:red; */
		}
	}

	/*************tablettttt******************/
	@media ${device.tablet} {
		#g-recaptcha {
			position: absolute;
			margin-left: 15%;
			bottom: 5%;
			transform: scale(0.7);
			transform-origin: 0 0;
		}
		max-width: 1000px;
		#hawk {
			position: absolute;
			z-index: -10;
			left: 10px;
			top: 10px;
			/* transform: translate(-50%,0%); */
			width: 7%;
		}
		#logo {
			position: absolute;
			z-index: -10;
			right: 10px;
			top: 16px;
			/* transform: translate(-50%,0%); */
			width: 5%;
		}
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
			width: 40%;
			position: absolute;
			height: 520px;
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
					.error {
						margin-left: auto;
						margin-right: auto;
						margin-top: 10px;
						text-align: center;
						width: 34%;
						padding: 5px;
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
		}

		#pcb1 {
			left: 50%;
			top: 0%;
			transform: translate(-50%, 0%);
			width: 100%;
			position: relative;
		}
		#pcb2 {
			left: 50%;
			bottom: 10px;
			transform: translate(-50%, 0%);
			width: 100%;
			position: relative;
		}
		#bg {
			animation: ${flash} 2s infinite 0s ease-in-out;
			opacity: 0.1;
			z-index: -100;
			position: absolute;
			overflow-y: hidden;
			left: 50%;
			top: 50%;
			transform: translate(-50%, -50%);
			width: 100vw;
			height: 100%;
			/* background:red; */
		}
	}

	/*************desktop******************/
	@media ${device.desktop} {
		#g-recaptcha {
			margin-bottom: -30px;
			margin-left: 25%;
			/* display: table; */
		}
		#pcbdesk {
			display: block;
			left: 50%;
			top: 0%;
			transform: translate(-50%, 0%);
			width: 100%;
			position: relative;
		}
		#pcb1,
		#pcb2 {
			display: none;
		}
		max-width: 3000px;
		#hawk {
			position: absolute;
			z-index: -10;
			left: 10px;
			top: 10px;
			/* transform: translate(-50%,0%); */
			width: 7%;
		}
		#logo {
			position: absolute;
			z-index: -10;
			right: 10px;
			top: 16px;
			/* transform: translate(-50%,0%); */
			width: 5%;
		}
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
			height: 520px;
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
					.error {
						margin-left: auto;
						margin-right: auto;
						margin-top: 10px;
						text-align: center;
						width: 34%;
						padding: 5px;
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
				cursor: pointer;
			}
		}

		#pcb1 {
			left: 50%;
			top: 0%;
			transform: translate(-50%, 0%);
			width: 100%;
			position: relative;
		}
		#pcb2 {
			left: 50%;
			bottom: 10px;
			transform: translate(-50%, 0%);
			width: 100%;
			position: relative;
		}
		#bg {
			animation: ${flash} 2s infinite 0s ease-in-out;
			opacity: 0.1;
			z-index: -100;
			position: absolute;
			overflow-y: hidden;
			left: 50%;
			top: 50%;
			transform: translate(-50%, -50%);
			width: 100vw;
			height: 100%;
			/* background:red; */
		}
	}
`;
