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
		loggedIn: false
	};

	onChange = (name, value) => {
		this.setState({ [name]: value });
	};

	login = () => {
		const { loggedIn, ...loginData } = this.state;
		fetch('/api/login', {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify(loginData)
		})
			.then(res => res.json())
			.then(json => {
				if (json.success) {
					this.setState({ loggedIn: true });
					this.props.UserStore.setCurrentUser(json.data);
				}
			});
	};

	render() {
		return (
			<div className={this.props.className}>
				<img src="hawk" alt="" id="hawk" />
				<div id="box">
					<h1>HawkEye</h1>
					<span>Sign into your account</span>
					<div id="inputs">
						<TextField
							name="username"
							placeholder="Username"
							onChange={this.onChange}
						/>
						<br />
						<TextField
							name="password"
							type="password"
							placeholder="Password"
							onChange={this.onChange}
						/>
						<br />
					</div>
					<Button onClick={this.login}>Sign In</Button>
					{this.state.loggedIn ? <Redirect to="/dashboard" /> : null}
					<br />
					<Link to="/register" id="register">
						Create an account
					</Link>
				</div>
				<img src={map} alt="" id="map" />
				<a href="https://www.iecsemanipal.com/">
					<img src={logo} alt="" id="logo" />
				</a>
			</div>
		);
	}
}

export default styled(LoginPage)`
	width: 100%;
	${media.phone`
	#box {
		${centerHV}
		padding-top:10px;
		border-radius:10px;
		background:#1c1c1c;
		filter: drop-shadow(0px 15px 15px #000);
		width:80%;
		position: absolute;
		height:280px;
		text-align:center;

		h1 {
			color:#FFD627;
			margin:10px 0 3px 0;
			text-transform: uppercase;
		}

		span {
			color:white;
			font-weight:300;
			font-size:0.7em;
			letter-spacing:3px;
		}

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
	#register{
		letter-spacing:1px;
		color:white;
		text-decoration:none;
		font-weight:300;
		font-size:0.8em;
	}

	}


	#map{
		position:absolute;
		z-index:-10;
		width:100%;
		top:-120px;
		left: 50%;
		opacity:0.25;
	transform: translate(-50%,0%);
	}
	#logo{
		position:absolute;
		width:10%;
		bottom:30px;
		left: 50%;
		transform: translate(-50%,0%);

	}
	`}
`;
