import React, { useCallback, useState } from 'react';
import PropTypes from 'prop-types';
import styled, { keyframes } from 'styled-components';
import validator from 'validator';
import media from '../components/theme/media';
import logo from '../components/assets/iecse_logo.png';
import hawk from '../components/assets/hawk_logo.png';
import pcb from '../components/assets/pcbdesign.png';
import pcb1 from '../components/assets/pcbdesign1.png';

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

const RegisterPage = ({ className }) => {
	const [formData, setformData] = useState({
		name: '',
		username: '',
		password: '',
		confirm_password: '',
		email: '',
		tel: '',
		college: ''
	});

	const onChange = useCallback((name, value, error) => {
		setformData(Object.assign(formData, { [name]: { value, error } }));
		console.log(formData);
	}, []);

	const onSubmit = useCallback(() => {}, []);

	return (
		<div className={className}>
			<div id="box">
				<h1>HAWKEYE</h1>
				<span>Create a new account</span>
				<div id="inputs">
					<TextField name="name" placeholder="Name" onChange={onChange} />
					<TextField
						name="username"
						placeholder="Username"
						onChange={onChange}
					/>
					<TextField
						name="password"
						type="password"
						placeholder="Password"
						onChange={onChange}
					/>
					<TextField
						name="confirm_password"
						type="password"
						placeholder="Confirm password"
						onChange={onChange}
					/>
					<TextField
						name="email"
						type="email"
						placeholder="Email"
						onChange={onChange}
						validation={v => (validator.isEmail(v) ? '' : 'Invalid Email')}
						validateOnChange
					/>
					<TextField
						name="tel"
						placeholder="Mobile Number"
						onChange={onChange}
						validation={v =>
							validator.isMobilePhone(v) ? '' : 'Invalid Number'
						}
						validateOnChange
					/>
					<TextField name="college" placeholder="College" onChange={onChange} />
				</div>
				<Button onClick={onSubmit}>Register</Button>
			</div>
			<img src={hawk} alt="" id="hawk" />
			<img src={logo} alt="" id="logo" />
			<div id="bg">
				<img src={pcb} alt="" id="pcb1" />
				<img src={pcb} alt="" id="pcb2" />
				<img src={pcb1} alt="" id="pcbdesk" />
			</div>
		</div>
	);
};

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
			height: 500px;
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
						text-align: center;
						width: 34%;
						padding: 5px;
						margin: 10px;
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
			height: 500px;
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
						text-align: center;
						width: 34%;
						padding: 5px;
						margin: 10px;
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
			height: 500px;
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
						text-align: center;
						width: 34%;
						padding: 5px;
						margin: 10px;
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
`;
