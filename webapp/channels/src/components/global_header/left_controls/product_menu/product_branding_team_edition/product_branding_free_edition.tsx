// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';
import {useSelector} from 'react-redux';
import styled from 'styled-components';

import {getLicense} from 'mattermost-redux/selectors/entities/general';

import Logo from 'components/common/svg_images_components/logo_dark_blue_svg';

import {LicenseSkus} from 'utils/constants';

const ProductBrandingFreeEditionContainer = styled.span`
    display: flex;
    align-items: center;
    overflow: visible;
    flex-shrink: 0;

    > * + * {
        margin-left: 10px;
    }
`;

const StyledLogo = styled(Logo)`
    opacity: 0.9;
`;

const Badge = styled.span`
    display: flex;
    align-self: center;
    padding: 4px 8px;
    border-radius: var(--radius-s);
    margin-left: 10px;
    background: rgba(var(--sidebar-text-rgb), 0.12);
    color: rgba(var(--sidebar-text-rgb), 0.9);
    font-family: 'Open Sans', sans-serif;
    font-size: 11px;
    font-weight: 700;
    letter-spacing: 0.05em;
    line-height: 16px;
    text-transform: uppercase;
`;

const ProductBrandingFreeEdition = (): JSX.Element => {
    const license = useSelector(getLicense);

    let badgeText = '';
    if (license?.SkuShortName === LicenseSkus.Entry) {
        badgeText = 'ENTRY EDITION';
    } else if (license?.IsLicensed === 'false') {
        badgeText = 'PING EDITION';
    }

    return (
        <ProductBrandingFreeEditionContainer tabIndex={-1}>
            <StyledLogo
                width={140}
                height={28}
            />
            {/* Removed: Team Edition badge */}
        </ProductBrandingFreeEditionContainer>
    );
};

export default ProductBrandingFreeEdition;
