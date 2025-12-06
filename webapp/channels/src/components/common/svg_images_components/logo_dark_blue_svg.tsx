// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';

import PingLogoImage from 'images/ping-logo.png';

type Props = {
    width?: number;
    height?: number;
    className?: string;
}

export default (props: Props) => (
    <img
        src={PingLogoImage}
        alt="Ping Logo"
        className={props.className}
        width={props.width}
        height={props.height}
        style={{maxWidth: props.width ? `${props.width}px` : '182px', maxHeight: props.height ? `${props.height}px` : '30px'}}
    />
);
