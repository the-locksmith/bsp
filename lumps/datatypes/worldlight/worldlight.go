package worldlight

import "github.com/galaco/bsp/lumps/datatypes/common"

/* Assuming this is 8bits
enum emittype_t
{
	emit_surface,		// 90 degree spotlight
	emit_point,			// simple point light source
	emit_spotlight,		// spotlight with penumbra
	emit_skylight,		// directional light with no falloff (surface must trace to SKY texture)
	emit_quakelight,	// linear falloff, non-lambertian
	emit_skyambient,	// spherical light source with no falloff (surface must trace to SKY texture)
};
 */

type WorldLight struct {
	Origin common.Vector
	Intensity common.Vector
	Normal common.Vector
	Cluster int32
	EmitType int8
	Style int32
	Stopdot float32
	Stopdot2 float32
	Exponent float32
	Radius float32
	ConstantAttenuation float32
	LinearAttenuation float32
	QuadraticAttenuation float32
	Flags int32
	TexInfo int32
	Owner int32
}