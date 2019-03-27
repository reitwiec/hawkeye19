import React, { useCallback, useState } from 'react';
import PropTypes from 'prop-types';
import styled, { keyframes } from 'styled-components';
import validator from 'validator';
import media from '../components/theme/media';
import logo from '../components/assets/iecse_logo.png';
import hawk from '../components/assets/hawk_logo.png';
import pcb from '../components/assets/pcbdesign.png';

import { Button, TextField } from '../components';

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
	}, []);

	const onSubmit = useCallback(() => {
		// Remove error keys from formData
		const postData = Object.entries(formData)
			.map(([k, v]) => {
				return { [k]: v.value };
			})
			.reduce((acc, val) => Object.assign(acc, val));

		fetch('/api/addUser', {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify(postData)
		})
			.then(res => res.json())
			.then(json => console.log(json));
	}, []);

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
	${media.phone`
	#hawk{
		position:absolute;
		z-index:-10;
		left:10px;
		top:10px;
		/* transform: translate(-50%,0%); */
		width:10%;
	}
	#logo{
		position:absolute;
		z-index:-10;
		right:10px;
		top:16px;
		/* transform: translate(-50%,0%); */
		width:8%;
	}
	#box{
		padding-top:10px;
		h1{
			color:#FFD627;
			margin:10px 0 3px 0;
		}
		span{
			color:white;
			font-weight:300;
			font-size:0.7em;
			letter-spacing:3px;
		}
		border-radius:10px;

		background:#1c1c1c;
		filter: drop-shadow(0px 15px 15px #000);
		width:80%;
		position: absolute;
		height:500px;
		text-align:center;
    left: 50%;
    top: 50%;
	transform: translate(-50%,-50%);
	#inputs{
		text-align:left;
		width:100%;
		margin-left:30px;
		${TextField}{
			input{
				letter-spacing:1.25px;
				display:block;
				font-size:100%;
				text-indent: 10px;
				margin-top:10px;
				font-weight:300;
				width:80%;
				color:#fff;
				border: 0px;
				border-radius: 4px;
				box-shadow: none;
				outline: none;
				-webkit-appearance:none ;
				-moz-appearance:none ;
				appearance:none ;
				height:25px;
				border:none;
				text-indent: 10px;
				background: #333333;
			}
			.error{
				text-align:center;
				width:34%;
				padding:5px;
				margin:10px;
			}
		}
	}
	${Button}{
		color: #1c1c1c;
		font-weight: 500;
		background: #FFD627;
		width: 50%;
		height: 35px;
		padding: 10px;
		padding-top: 7px;
		border: none;
		border-radius: 20px;
		margin-top:10px;
		margin-bottom:10px;

	}

	}


	#pcb1{
		left: 50%;
    top: 0%;
	transform: translate(-50%,0%);
	width:100%;
		position:relative;
		
	}
	#pcb2{
		left: 50%;
    bottom: 10px;
	transform: translate(-50%,0%);
	width:100%;
		position:relative;

	}
	#bg{
		animation: ${flash} 2s infinite 0s ease-in-out;
		opacity:0.1;
		z-index:-100;
		position:absolute;
		overflow-y: hidden;
		left: 50%;
    top: 50%;
	transform: translate(-50%,-50%);
	width: 100vw;
	
	height: 100%;
	/* background:red; */
	}
`}
`;
