// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';
import styled from 'styled-components';

import Logo from 'components/common/svg_images_components/logo_dark_blue_svg';

const ProductBrandingFreeEditionContainer = styled.span`
    display: flex;
    align-items: center;
    overflow: visible;
    flex-shrink: 0;
`;

const StyledLogo = styled(Logo)`
    opacity: 0.9;
`;

const ProductBrandingFreeEdition = (): JSX.Element => {
    return (
        <ProductBrandingFreeEditionContainer tabIndex={-1}>
            <StyledLogo
                width={140}
                height={28}
            />
        </ProductBrandingFreeEditionContainer>
    );
};

export default ProductBrandingFreeEdition;
