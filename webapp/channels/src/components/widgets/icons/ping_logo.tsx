// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';
import {useIntl} from 'react-intl';

import PingLogoImage from 'images/ping-logo.png';

export default function PingLogo(props: React.HTMLAttributes<HTMLSpanElement>) {
    const {formatMessage} = useIntl();
    return (
        <span {...props}>
            <img
                src={PingLogoImage}
                alt={formatMessage({id: 'generic_icons.ping', defaultMessage: 'Ping Logo'})}
                role='img'
                aria-label={formatMessage({id: 'generic_icons.ping', defaultMessage: 'Ping Logo'})}
            />
        </span>
    );
}
