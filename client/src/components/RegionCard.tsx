import React from 'react';
import styled from 'styled-components';
import { Attributes } from '../utils';
import { Link } from 'react-router-dom';
import colors from '../colors';
import { square } from '../mixins';

interface IRegionCardProps extends Attributes<'div'> {
	name: string;
	icon: string;
	locked: boolean;
}

const RegionCard: React.SFC<IRegionCardProps> = props => {
	const { name, icon, ...restProps } = props;
	return (
		<RegionCardWrapper className="region-card" {...restProps}>
			<StyledLink to={{ pathname: '/questions', state: { name } }}>
				<img className="icon" src={icon} alt="" />
				<span className="name">{name}</span>
			</StyledLink>
		</RegionCardWrapper>
	);
};

const StyledLink = styled(Link)`
	color: white;
	text-decoration: none;
	text-shadow: 1px 1px 2px black;
`;

const RegionCardWrapper = styled.div`
	cursor: pointer;
	background-color: ${colors.mediumGrey};
	border-radius: 100%;
	text-align: center;
	padding: 5px;
`;

export default RegionCard;
