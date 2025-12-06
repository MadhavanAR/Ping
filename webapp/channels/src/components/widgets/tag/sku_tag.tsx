// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import classNames from 'classnames';
import React, {useMemo} from 'react';

import {LicenseSkus} from 'utils/constants';

import Tag from './tag';
import type {TagSize} from './tag';

type Props = {
    className?: string;
    size?: TagSize;
    sku: LicenseSkus;
};

const SkuTag = ({className = '', size = 'xs', sku}: Props) => {
    const namedSku = useMemo(() => {
        switch (sku) {
        case LicenseSkus.Starter:
            return 'STARTER';
        case LicenseSkus.Professional:
            return 'PROFESSIONAL';
        case LicenseSkus.Enterprise:
            return 'PING';
        case LicenseSkus.E10:
            return 'PING';
        case LicenseSkus.E20:
            return 'PING';
        case LicenseSkus.EnterpriseAdvanced:
            return 'PING';
        case LicenseSkus.Entry:
            return 'PING';
        default:
            return 'PING';
        }
    }, [sku]);

    return (
        <Tag
            className={classNames('SkuTag', className)}
            icon='ping'
            size={size}
            text={namedSku}
        />
    );
};

export default SkuTag;
