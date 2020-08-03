// +build ios

/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2019 WireGuard LLC. All Rights Reserved.
 */

package device

/* Fit within memory limits for iOS's Network Extension API, which has stricter requirements */

const (
	QueueOutboundSize          = 1024
	QueueInboundSize           = 1024
	QueueHandshakeSize         = 1024
	MaxSegmentSize             = 1700
	// spalib会额外占用一些内存，导致按原来的长度申请缓冲会超过ios Extension的内存限制，由1024调整为512
	PreallocatedBuffersPerPool = 512
)
