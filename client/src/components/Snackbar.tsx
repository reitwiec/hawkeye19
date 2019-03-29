import React, { Component } from 'react';
import styled, { css } from 'styled-components';

interface ISnackbarProps {
	message: JSX.Element | string;
	action: JSX.Element | null;
	open: boolean;
	autoHideDelay: number;
}

class Snackbar extends Component<ISnackbarProps> {
	public static defaultProps = {
		action: null,
		autoHideDelay: 0,
		open: false
	};

	public render() {
		return (
			<SnackbarWrapper {...this.props}>{this.props.message}</SnackbarWrapper>
		);
	}
}

const SnackbarWrapper = styled.div`
	display: block;
	position: fixed;
	top: 0;
	width: 100vw;
	background-color: #ff0000;
	padding: 0.5em;
	text-align: center;
	color: #fff;
	text-transform: uppercase;
	opacity: 0;
	transition: all 0.25s ease;
	margin-bottom: -2em;
	z-index: 999;
	font-weight: 500;
	letter-spacing: 1.5px;
	filter: drop-shadow(0px 15px 15px #000);

	@media only screen and (max-width: 500px) {
		font-size: 0.8rem;
	}

	${(props: ISnackbarProps) =>
		props.open
			? css`
					opacity: 1;
					margin-bottom: 0;
			  `
			: ``}
`;

export default Snackbar;
