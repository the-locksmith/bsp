package bsp

type LumpId int

const (
	LumpEntities 					= LumpId(0)   // Entity keyvalue data stored as string
	LumpPlanes 						= LumpId(1)     // bsp planes
	LumpTexData 					= LumpId(2)    // texture data used by bsp faces
	LumpVertexes 					= LumpId(3)   // vertex data
	LumpVisibility 					= LumpId(4) // vvis calculated visibility pvs & pas information
	LumpNodes 						= LumpId(5)      // bsp node tree entries
	LumpTexInfo 					= LumpId(6)    // face texture information
	LumpFaces 						= LumpId(7)      // bsp faces
	LumpLighting 					= LumpId(8)
	LumpOcclusion 					= LumpId(9)
	LumpLeafs 						= LumpId(10)
	LumpFaceIds 					= LumpId(11 )  // contents is normally stripped out by compiler
	LumpEdges 						= LumpId(12)     // face edges. v1->v2, vertex order may be reversed
	LumpSurfEdges 					= LumpId(13) //
	LumpModels 						= LumpId(14)    // models are root bsp nodes. m[0] = worldspawn. m[0+n] are brush entity data
	LumpWorldLights 				= LumpId(15)
	LumpLeafFaces 					= LumpId(16 )  // faces that separate leaves.
	LumpLeafBrushes 				= LumpId(17) // brushes that define a leaf volume
	LumpBrushes 					= LumpId(18)
	LumpBrushSides 					= LumpId(19)
	LumpAreas 						= LumpId(20)
	LumpAreaPortals 				= LumpId(21)
	LumpPortals 					= LumpId(22)
	LumpUnused0 					= LumpId(22)
	LumpPropCollision 				= LumpId(22)
	LumpClusters 					= LumpId(23)
	LumpUnused1 					= LumpId(23)
	LumpPropHulls					= LumpId(23)
	LumpPortalVerts 				= LumpId(24)
	LumpUnused2 					= LumpId(24)
	LumpPropHullVerts 				= LumpId(24)
	LumpClusterPortals 				= LumpId(25)
	LumpUnused3 					= LumpId(25)
	LumpPropTris 					= LumpId(25)
	LumpDispInfo 					= LumpId(26)
	LumpOriginalFaces 				= LumpId(27)
	LumpPhysDisp 					= LumpId(28)
	LumpPhysCollide 				= LumpId(29)
	LumpVertNormals 				= LumpId(30)
	LumpVertNormalIndices 			= LumpId(31)
	LumpDispLightmapAlphas 			= LumpId(32) // contents is normally stripped out
	LumpDispVerts 					= LumpId(33)
	LumpDispLightmapSamplePositions = LumpId(34)
	LumpGame 						= LumpId(35) // game specific data. includes staticprop data
	LumpLeafWaterData 				= LumpId(36)
	LumpPrimitives					= LumpId(37)
	LumpPrimVerts 					= LumpId(38)
	LumpPrimIndices 				= LumpId(39)
	LumpPakfile 					= LumpId(40) // uncompressed zip of packed custom content
	LumpClipPortalVerts 			= LumpId(41)
	LumpCubemaps 					= LumpId(42)
	LumpTexDataStringData 			= LumpId(43)  // raw string data of material paths
	LumpTexDataStringTable			= LumpId(44) // references entries in the string data lump
	LumpOverlays					= LumpId(45)
	LumpLeafMinDistToWater 			= LumpId(46)
	LumpFaceMacroTextureInfo 		= LumpId(47)
	LumpDispTris 					= LumpId(48)
	LumpPhysCollideSurface 			= LumpId(49) // deprecated
	LumpPropBlob 					= LumpId(49)
	LumpWaterOverlays 				= LumpId(50)
	LumpLightmapPages 				= LumpId(51)
	LumpLeafAmbientIndexHDR 		= LumpId(51)
	LumpLightmapPagesInfo 			= LumpId(52)
	LumpLeafAmbientIndex 			= LumpId(52)
	LumpLightingHDR 				= LumpId(53)
	LumpWorldLightsHDR 				= LumpId(54)
	LumpLeafAmbientLightingHDR 		= LumpId(55)
	LumpLeafAmbientLighting 		= LumpId(56)
	LumpXZipPakfile 				= LumpId(57) // deprecated, and xbox specific
	LumpFacesHDR 					= LumpId(58)
	LumpMapFlags 					= LumpId(59)
	LumpOverlayFades				= LumpId(60)
	LumpOverlaySystemLevels 		= LumpId(61)
	LumpPhysLevel 					= LumpId(62)
	LumpDispMultiBlend 				= LumpId(63)
)