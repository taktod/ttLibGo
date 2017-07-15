#ifndef TTLIBC_LOG_H_
#define TTLIBC_LOG_H_

#include <stdio.h>

#ifdef __cplusplus
extern "C" {
#endif

#ifndef __DEBUG_FLAG__
#	define __DEBUG_FLAG__ 1
#endif

/**
 * output log data. only for debug compile
 * @param fmt format string
 * @param ... data for format string
 */
#if __DEBUG_FLAG__ == 1
#	define	LOG_PRINT(fmt, ...) \
	printf("[log]%s(): " fmt "\n", __func__, ## __VA_ARGS__)
#else
#	define	LOG_PRINT(fmt, ...)
#endif

/**
 * output error data for strerr.
 * @param fmt format string
 * @param ... data for format string
 */
#define	ERR_PRINT(fmt, ...) \
	printf("[log]%s(): " fmt "\n", __func__, ## __VA_ARGS__)

#ifdef __cplusplus
} /* extern "C" */
#endif

#endif
